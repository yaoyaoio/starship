package module

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/tricorder/src/api-server/dao"
	modulepb "github.com/tricorder/src/pb/module"
	"github.com/tricorder/src/utils/file"
	"github.com/tricorder/src/utils/uuid"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init module by json file.",
	Long: `init module by json file. For example:
$ starship-cli module init -b path/to/bcc_file -m path/to/module_json_file -w path/to/wasm_file`,
	Run: func(cmd *cobra.Command, args []string) {
		bccStr, err := file.Read(bccFilePath)
		if err != nil {
			log.Error(err)
		}

		wasmBytes, err := file.ReadBin(wasmFilePath)
		if err != nil {
			log.Error(err)
		}

		moduleReq, err := parseModuleJsonFile(moduleFilePath)
		if err != nil {
			log.Error(err)
		}
		// override bcc code contet by bcc file
		moduleReq.Ebpf.Code = bccStr
		// override wasm code contet by wasm file
		moduleReq.Wasm.Code = wasmBytes
		err = initModule(moduleReq)
		if err != nil {
			log.Error(err)
		}
	},
}

var dbFilePath string

func init() {
	initCmd.Flags().StringVarP(&moduleFilePath, "module-file-path", "m",
		moduleFilePath, "The file path of module in json format.")
	initCmd.Flags().StringVarP(&bccFilePath, "bcc-file-path", "b", bccFilePath, "The file path of bcc code.")
	initCmd.Flags().StringVarP(&wasmFilePath, "wasm-file-path", "w", wasmFilePath, "The file path of wasm code.")
	initCmd.Flags().StringVarP(&dbFilePath, "db-file-path", "d", dbFilePath, "The file path of sqlit db.")
}

func initModule(body *modulepb.Module) error {
	sqliteClient, _ := dao.InitSqlite(dbFilePath)
	codeDao := dao.Module{
		Client: sqliteClient,
	}

	ebpfProbes, err := json.Marshal(body.Ebpf.Probes)
	if err != nil {
		return err
	}

	schemaAttr, err := json.Marshal(body.Wasm.OutputSchema.Fields)
	if err != nil {
		return err
	}

	mod := &dao.ModuleGORM{
		ID:                 strings.Replace(uuid.New(), "-", "_", -1),
		Name:               body.Name,
		CreateTime:         time.Now().Format("2006-01-02 15:04:05"),
		Status:             int(0),
		Ebpf:               body.Ebpf.Code,
		EbpfFmt:            int(body.Ebpf.Fmt),
		EbpfLang:           int(body.Ebpf.Lang),
		EbpfPerfBufferName: body.Ebpf.PerfBufferName,
		EbpfProbes:         string(ebpfProbes),
		Wasm:               body.Wasm.Code,
		SchemaAttr:         string(schemaAttr),
		Fn:                 body.Wasm.FnName,
		WasmFmt:            int(body.Wasm.Fmt),
		WasmLang:           int(body.Wasm.Lang),
	}

	mod.SchemaName = fmt.Sprintf("%s_%s", "tricorder_code", mod.ID)

	err = codeDao.SaveCode(mod)
	return err
}

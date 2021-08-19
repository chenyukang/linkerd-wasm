package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/second-state/WasmEdge-go/wasmedge"
)

func imageHandlerWASI(w http.ResponseWriter, r *http.Request) {
	image, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	/// Set not to print debug info
	wasmedge.SetLogErrorLevel()

	/// Create configure
	var conf = wasmedge.NewConfigure(wasmedge.REFERENCE_TYPES)
	conf.AddConfig(wasmedge.WASI)

	/// Create VM with configure
	var vm = wasmedge.NewVMWithConfig(conf)

	/// Init WASI (test)
	var wasi = vm.GetImportObject(wasmedge.WASI)
	wasi.InitWasi(
		os.Args[1:],     /// The args
		os.Environ(),    /// The envs
		[]string{".:."}, /// The mapping directories
		[]string{},      /// The preopens will be empty
	)

	/// Register WasmEdge-tensorflow and WasmEdge-image
	var tfobj = wasmedge.NewTensorflowImportObject()
	var tfliteobj = wasmedge.NewTensorflowLiteImportObject()
	vm.RegisterImport(tfobj)
	vm.RegisterImport(tfliteobj)
	var imgobj = wasmedge.NewImageImportObject()
	vm.RegisterImport(imgobj)
	/// Instantiate wasm

	vm.LoadWasmFile("./lib/classify_lib_bg.wasm")
	vm.Validate()
	vm.Instantiate()

	res, err := vm.ExecuteBindgen("infer", wasmedge.Bindgen_return_array, image)
	ans := string(res.([]byte))
	if err != nil {
		println("error: ", err.Error())
	}

	vm.Delete()
	conf.Delete()

	fmt.Printf("Image classify result: %q\n", ans)
	fmt.Fprintf(w, "%s", ans)
}

func main() {
	http.HandleFunc("/api", imageHandlerWASI)
	println("listen to 8086 ...")
	log.Fatal(http.ListenAndServe(":8086", nil))
}

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/core"
	"github.com/IBM/whcs-go-sdk/annotatorforclinicaldataacdv1"
)

var pipelineOptions *annotatorforclinicaldataacdv1.RunPipelineOptions
var ACD *annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1

func init() {
	apikey := os.Getenv("WH_ACD_APIKEY")
	version := "2020-11-02"
	url := os.Getenv("WH_ACD_URL")

	var err error

	ACD, err = annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
		URL:     url,
		Version: core.StringPtr(version),
		Authenticator: &core.IamAuthenticator{
			ApiKey: apikey,
		},
	})
	if err != nil {
		panic(err)
	}

	pipelineOptions = ACD.NewRunPipelineOptions()
	cdParams := make(map[string][]string)
	cdParamValue := []string{"false"}
	cdParams["apply_spell_check"] = cdParamValue
	cdAnnotator, err := ACD.NewAnnotator("symptom_disease")
	cdAnnotator.Parameters = cdParams
	cdFlowEntry, err := ACD.NewFlowEntry(cdAnnotator)
	cdFlowEntry.Annotator = cdAnnotator
	async := false
	flow, err := ACD.NewFlow([]annotatorforclinicaldataacdv1.FlowEntry{*cdFlowEntry}, core.BoolPtr(async))
	annotatorFlow, err := ACD.NewAnnotatorFlow(flow)
	pipelineOptions.SetAnnotatorFlows([]annotatorforclinicaldataacdv1.AnnotatorFlow{*annotatorFlow})
}

func Annotator(diagnose string) []interface{} {

	container := ACD.NewUnstructuredContainer()
	container.SetText(diagnose)
	pipelineOptions.SetUnstructured([]annotatorforclinicaldataacdv1.UnstructuredContainer{*container})
	pipelineOptions.SetDebugTextRestore(false)

	_, detailResponse, err := ACD.RunPipeline(pipelineOptions)
	if err != nil {
		fmt.Println("Error in run pipeline")
		panic(err)
	}
	fmt.Println(detailResponse.String())

	// fmt.Println(detailResponse.GetResult())
	// fmt.Println(&result.Unstructured)

	var result map[string]interface{}
	err = json.Unmarshal([]byte(detailResponse.String()), &result)
	if err != nil {
		panic(err)
	}
	mapResult := result["Result"].(map[string]interface{})
	unstruct := mapResult["unstructured"].([]interface{})
	datas := unstruct[0].(map[string]interface{})
	symptoms := datas["data"].(map[string]interface{})
	icds := symptoms["SymptomDiseaseInd"].([]interface{})

	// output, err := json.MarshalIndent(icds, "", " ")

	// diagnoseResult := strings.ReplaceAll(strings.ReplaceAll(string(output), "\n", ""), "\\", "")

	// fmt.Println(diagnoseResult)

	return icds
}
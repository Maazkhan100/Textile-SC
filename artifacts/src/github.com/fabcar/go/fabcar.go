package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type TextileChaincode struct {
	contractapi.Contract
}

type BatchHistory struct {
    BatchID                  	string                   	`json:"batchId"`
    CurrentStage             	string                   	`json:"currentStage"`
    FiberProducerBatch       	FiberProducerData   		`json:"fiberProducerBatch"`
    YarnManufacturerBatch   	YarnManufacturerData 		`json:"yarnManufacturerBatch"`
	FabricManufacturerBatch		FabricManufacturerData		`json:"fabricManufacturerBatch"`
	ApparelManufacturerBatch	ApparelManufacturerData		`json:"apparelManufacturerBatch"`
	RetailerBatch				RetailerData				`json:"retailerBatch"`
}

type FiberProducerBasicData struct {
    FiberProducerID     string					`json:"fiberProducerID"`
    Name                string					`json:"name"`
    Address				string					`json:"address"`
    FiberProducer		[]FiberProducerData		`json:"fiberProducer"`
}

type FiberProducerData struct {
	BatchID                  		string                 `json:"batchId"`
    ProducerName     		 		string                 `json:"producerName"`
    FiberType         		 		string                 `json:"fiberType"`
    Location           		 		string                 `json:"location"`
    HarvestDate        		 		string                 `json:"harvestDate"`
    LandArea           		 		float64                `json:"landArea"`
    ProductionCapacity 		 		int                    `json:"productionCapacity"`
    FarmingPractices    	 		string                 `json:"farmingPractices"`
    SoilHealth         		 		string                 `json:"soilHealth"`
    Yield              		 		int                    `json:"yield"`
    FertilizerName     	 	 		string                 `json:"fertilizerName"`
    FertilizerQuantity 		 		int                    `json:"fertilizerQuantity"`
    PesticideName      		 		string                 `json:"pesticideName"`
    PesticideQuantity  		 		int                    `json:"pesticideQuantity"`
    HerbicideName      		 		string                 `json:"herbicideName"`
    HerbicideQuantity  		 		int                    `json:"herbicideQuantity"`
    WaterUsage         		 		float64                `json:"waterUsage"`
    Inventory          		 		float64                `json:"inventory"`
	FiberProducingDate		 		string 					`json:"fiberProducingDate"`
    FiberProducerTransactions       []TransactionsList 		`json:"fiberProducerTransactions"`
}

type TransactionsList struct {
	TransactionDate				string				`json:"transactionDate"`
	TransactionRecipient		string				`json:"transactionRecipient"`
	TransactionQuantity			float64				`json:"transactionQuantity"`
	TransactionFiberType		string				`json:"transactionFiberType"`
}

type YarnManufacturerBasicData struct {
    YarnManufacturerID    string						`json:"yarnManufacturerID"`
    Name                  string 						`json:"name"`
    Address               string 					 	`json:"address"`
    YarnManufacturer	  []YarnManufacturerData 		`json:"yarnManufacturer"`
}

type YarnManufacturerData struct {
	BatchID                  			string                   `json:"batchId"`
    CompanyName        					string                   `json:"companyName"`
    FiberTypeUsed      					string                   `json:"fiberTypeUsed"`
    FiberSource        					string                   `json:"fiberSource"`
    YarnTypeProduced   					string                   `json:"yarnTypeProduced"`
    YarnCount          					string                   `json:"yarnCount"`
    ProductionCapacity 					int                      `json:"productionCapacity"`
    Dyeing             					string                   `json:"dyeing"`
    WasteManagement    					float64                  `json:"wasteManagement"`
    FinishingTreatment 					string                   `json:"finishingTreatment"`
    Strength           					string                   `json:"strength"`
    ColorsUsed         					[]string                 `json:"colorsUsed"`
    TwistLevel         					string                   `json:"twistLevel"`
    EnergyConsumption  					float64                  `json:"energyConsumption"`
    Machinery          					string                   `json:"machinery"`
    Inventory          					float64                  `json:"inventory"`
	YarnManufacturingDate				string 					 `json:"yarnManufacturingDate"`
    YarnManufacturerTransactions        []TransactionsList 		 `json:"yarnManufacturerTransactions"`
}

type FabricManufacturerBasicData struct {
    FabricManufacturerID  string 					  `json:"fabricManufacturerID"`
    Name                  string 					  `json:"name"`
    Address               string   					  `json:"address"`
    FabricManufacturer    []FabricManufacturerData    `json:"fabricManufacturer"`
}

type FabricManufacturerData struct {
	BatchID                  		string                   	`json:"batchId"`
    CompanyName          			string   					`json:"companyName"`
    YarnTypeUsed         			string   					`json:"yarnTypeUsed"`
    FabricProducedType   			string   					`json:"fabricProducedType"`
    FabricComposition    			[]FiberComposition 			`json:"fabricComposition"`
    FinishingTreatment   			[]string 					`json:"finishingTreatment"`
    DefectRatio          			float64  					`json:"defectRatio"`
    FabricColor          			string   					`json:"fabricColor"`
    FabricTexture        			string   					`json:"fabricTexture"`
    MachineSpecification 			string   					`json:"machineSpecification"`
    FabricWaterConsumption 			float64 					`json:"fabricWaterConsumption"`
    ProductionTechniques 			[]string 					`json:"productionTechniques"`
    FabricManufacturingDate 		string 						`json:"fabricManufacturingDate"`
    Inventory            			float64  					`json:"inventory"`
	FabricProducerTransactions 		[]TransactionsList 			`json:"fabricProducerTransactions"`
}

type FiberComposition struct {
    FiberType   string  `json:"fiberType"`
    Percentage  float64 `json:"percentage"`
}

type ApparelManufacturerBasicData struct {
    ApparelManufacturerID       string 							`json:"id"`
    Name                  		string 							`json:"name"`
    Address               		string 							`json:"address"`
    ApparelManufacturer   		[]ApparelManufacturerData  		`json:"apparelManufacturer"`
}

type ApparelManufacturerData struct {
	BatchID                  				string                   	`json:"batchId"`
    CompanyName          					string   					`json:"companyName"`
    FabricTypeUsed       					string   					`json:"fabricTypeUsed"`
    GarmentTypeProduced  					string   					`json:"garmentTypeProduced"`
    ProductionVolume     					int      					`json:"productionVolume"`
    ManufacturingProcess 					[]string 					`json:"manufacturingProcess"`
    MaterialWaste        					float64  					`json:"materialWaste"`
    SizeRange            					string   					`json:"sizeRange"`
    Color                					[]string 					`json:"color"`
    StyleDetail          					string   					`json:"styleDetail"`
    ApparelMachineSpecification 			string 						`json:"apparelMachineSpecification"`
    Inventory            					float64  					`json:"inventory"`
    ApparelManufacturingDate				string						`json:"apparelManufacturingDate"`
    ApparelManufacturerTransactions         []TransactionsList 			`json:"apparelManufacturerTransactions"`
}

type RetailerBasicData struct {
    RetailerID      string 				`json:"id"`
    Name          	string 				`json:"name"`
    Address       	string 				`json:"address"`
    RetailerData  	[]RetailerData 		`json:"retailerData"`
}

type RetailerData struct {
	BatchID                 string                   	`json:"batchId"`
    StoreName         		string 						`json:"storeName"`
    StoreAddress      		string 						`json:"storeAddress"`
    ProductAssortment 		[]string 					`json:"productAssortment"`
    Inventory         		[]GarmentInventory 			`json:"inventory"`
}

type GarmentInventory struct {
    Type     		string 		`json:"type"`
    Color    		string 		`json:"color"`
    Brand    		string 		`json:"brand"`
    Size     		string 		`json:"size"`
    Quantity 		int    		`json:"quantity"`
}

// Init initializes the chaincode
func (cc *TextileChaincode) Init(ctx contractapi.TransactionContextInterface) error {
	// Initialize the ledger for all stakeholders
	err := cc.initFiberProducers(ctx)
	if err != nil {
		return fmt.Errorf("error initializing fiber producers: %v", err)
	}

	err = cc.initYarnManufacturers(ctx)
	if err != nil {
		return fmt.Errorf("error initializing yarn manufacturers: %v", err)
	}

	err = cc.initFabricManufacturers(ctx)
	if err != nil {
		return fmt.Errorf("error initializing fabric manufacturers: %v", err)
	}

	err = cc.initApparelManufacturers(ctx)
	if err != nil {
		return fmt.Errorf("error initializing apparel manufacturers: %v", err)
	}

	err = cc.initRetailers(ctx)
	if err != nil {
		return fmt.Errorf("error initializing retailers: %v", err)
	}

	return nil
}

// initFiberProducers initializes the ledger for a fiber producer
func (cc *TextileChaincode) initFiberProducers(ctx contractapi.TransactionContextInterface) error {
	// Initialize FiberProducerTransactionsList
	fiberProducerTransaction := TransactionsList{
		TransactionDate:         "2024-03-28",
		TransactionRecipient:    "YM1",
		TransactionQuantity:     100.0,
		TransactionFiberType:    "Cotton",
	}

	// Marshal fiber producer transaction data
	fiberProducerTransactionJSON, err := json.Marshal(fiberProducerTransaction)
	if err != nil {
		return fmt.Errorf("failed to marshal fiber producer transaction data: %v", err)
	}

	// Put fiber producer transaction data on ledger
	err = ctx.GetStub().PutState("fiberProducerTransaction", fiberProducerTransactionJSON)
	if err != nil {
		return fmt.Errorf("failed to put fiber producer transaction data on ledger: %v", err)
	}

	// Initialize FiberProducerData
	fiberProducerData := FiberProducerData{
		ProducerName:               "Producer 1",
		FiberType:                  "Type 1",
		Location:                   "Location 1",
		HarvestDate:                "2022-01-01",
		LandArea:                   100.5,
		ProductionCapacity:         500,
		FarmingPractices:           "Practices 1",
		SoilHealth:                 "Good",
		Yield:                      200,
		FertilizerName:             "Fertilizer 1",
		FertilizerQuantity:         100,
		PesticideName:              "Pesticide 1",
		PesticideQuantity:          50,
		HerbicideName:              "Herbicide 1",
		HerbicideQuantity:          30,
		WaterUsage:                 300.5,
		Inventory:                  1000,
		FiberProducingDate:         "2022-01-01",
		FiberProducerTransactions:  []TransactionsList{fiberProducerTransaction},
	}

	// Marshal fiber producer data
	fiberProducerDataJSON, err := json.Marshal(fiberProducerData)
	if err != nil {
		return fmt.Errorf("failed to marshal fiber producer data: %v", err)
	}

	// Put fiber producer data on ledger
	err = ctx.GetStub().PutState("1", fiberProducerDataJSON)
	if err != nil {
		return fmt.Errorf("failed to put fiber producer data on ledger: %v", err)
	}

	// Initialize FiberProducerBasicData
	fiberProducerBasicData := FiberProducerBasicData{
		FiberProducerID: "1",
		Name:            "Fiber Producer 1",
		Address:         "Address 1",
		FiberProducer:   []FiberProducerData{fiberProducerData},
	}

	// Marshal fiber producer basic data
	fiberProducerBasicDataJSON, err := json.Marshal(fiberProducerBasicData)
	if err != nil {
		return fmt.Errorf("failed to marshal fiber producer basic data: %v", err)
	}

	// Put fiber producer basic data on ledger
	err = ctx.GetStub().PutState(fiberProducerBasicData.FiberProducerID, fiberProducerBasicDataJSON)
	if err != nil {
		return fmt.Errorf("failed to put fiber producer basic data on ledger: %v", err)
	}

	return nil
}

// initYarnManufacturers initializes the ledger for a yarn manufacturer
func (cc *TextileChaincode) initYarnManufacturers(ctx contractapi.TransactionContextInterface) error {
	// Initialize YarnManufacturerTransactionsList
	yarnManufacturerTransaction := TransactionsList{
		TransactionDate:      "2024-03-28",
		TransactionRecipient: "FM1",
		TransactionQuantity: 200.0,
		TransactionFiberType: "Cotton",
	}

	// Initialize YarnManufacturerData
	yarnManufacturerData := YarnManufacturerData{
		CompanyName:                 "Yarn Manufacturer 1",
		FiberTypeUsed:               "Cotton",
		FiberSource:                 "Source 1",
		YarnTypeProduced:            "Type 1",
		YarnCount:                   "Count 1",
		ProductionCapacity:          1000,
		Dyeing:                      "Dyeing 1",
		WasteManagement:             20.5,
		FinishingTreatment:          "Treatment 1",
		Strength:                    "Strong",
		ColorsUsed:                  []string{"Red", "Blue", "Green"},
		TwistLevel:                  "High",
		EnergyConsumption:           500.5,
		Machinery:                   "Machinery 1",
		Inventory:                   2000,
		YarnManufacturingDate:       "2022-01-01",
		YarnManufacturerTransactions: []TransactionsList{yarnManufacturerTransaction},
	}

	// Marshal yarn manufacturer data
	yarnManufacturerDataJSON, err := json.Marshal(yarnManufacturerData)
	if err != nil {
		return fmt.Errorf("failed to marshal yarn manufacturer data: %v", err)
	}

	// Put yarn manufacturer data on ledger
	err = ctx.GetStub().PutState("1", yarnManufacturerDataJSON)
	if err != nil {
		return fmt.Errorf("failed to put yarn manufacturer data on ledger: %v", err)
	}

	// Initialize YarnManufacturerBasicData
	yarnManufacturerBasicData := YarnManufacturerBasicData{
		YarnManufacturerID: "1",
		Name:               "Yarn Manufacturer 1",
		Address:            "Address 1",
		YarnManufacturer:  []YarnManufacturerData{yarnManufacturerData},
	}

	// Marshal yarn manufacturer basic data
	yarnManufacturerBasicDataJSON, err := json.Marshal(yarnManufacturerBasicData)
	if err != nil {
		return fmt.Errorf("failed to marshal yarn manufacturer basic data: %v", err)
	}

	// Put yarn manufacturer basic data on ledger
	err = ctx.GetStub().PutState(yarnManufacturerBasicData.YarnManufacturerID, yarnManufacturerBasicDataJSON)
	if err != nil {
		return fmt.Errorf("failed to put yarn manufacturer basic data on ledger: %v", err)
	}

	return nil
}

func (cc *TextileChaincode) initFabricManufacturers(ctx contractapi.TransactionContextInterface) error {
	// Initialize FabricComposition
	fabricComposition := []FiberComposition{
		{FiberType: "Cotton", Percentage: 70},
		{FiberType: "Polyester", Percentage: 30},
	}

	fabricCompositionJSON, err := json.Marshal(fabricComposition)
	if err != nil {
		return fmt.Errorf("failed to marshal fabric composition data: %v", err)
	}

	// Put fabricCompositionData on ledger
	err = ctx.GetStub().PutState("1", fabricCompositionJSON)
	if err != nil {
		return fmt.Errorf("failed to put fabric composition data on ledger: %v", err)
	}

	// Initialize TransactionsList
	fabricProducerTransaction := TransactionsList{
		TransactionDate:        "2024-03-28",
		TransactionRecipient:   "AM1",
		TransactionQuantity:    300.0,
		TransactionFiberType:   "Cotton",
	}

	// Initialize FabricManufacturerData
	fabricManufacturerData := FabricManufacturerData{
		CompanyName:          "Fabric Manufacturer 1",
		YarnTypeUsed:         "Cotton",
		FabricProducedType:   "Woven",
		FabricComposition:    fabricComposition,
		FinishingTreatment:   []string{"Wrinkle resistant", "Water resistant"},
		DefectRatio:          0.05,
		FabricColor:          "White",
		FabricTexture:        "Soft",
		MachineSpecification: "Specification 1",
		FabricWaterConsumption: 150.5,
		ProductionTechniques: []string{"Weaving"},
		FabricManufacturingDate: "2022-01-01",
		Inventory:            3000,
		FabricProducerTransactions: []TransactionsList{fabricProducerTransaction},
	}

	// Marshal FabricManufacturerData
	fabricManufacturerDataJSON, err := json.Marshal(fabricManufacturerData)
	if err != nil {
		return fmt.Errorf("failed to marshal fabric manufacturer data: %v", err)
	}

	// Put FabricManufacturerData on ledger
	err = ctx.GetStub().PutState("1", fabricManufacturerDataJSON)
	if err != nil {
		return fmt.Errorf("failed to put fabric manufacturer data on ledger: %v", err)
	}

	// Initialize FabricManufacturerBasicData
	fabricManufacturerBasicData := FabricManufacturerBasicData{
		FabricManufacturerID:                  "1",
		Name:                "Fabric Manufacturer 1",
		Address:             "Address 1",
		FabricManufacturer:  []FabricManufacturerData{fabricManufacturerData},
	}

	// Marshal FabricManufacturerBasicData
	fabricManufacturerBasicDataJSON, err := json.Marshal(fabricManufacturerBasicData)
	if err != nil {
		return fmt.Errorf("failed to marshal fabric manufacturer basic data: %v", err)
	}

	// Put FabricManufacturerBasicData on ledger
	err = ctx.GetStub().PutState(fabricManufacturerBasicData.FabricManufacturerID, fabricManufacturerBasicDataJSON)
	if err != nil {
		return fmt.Errorf("failed to put fabric manufacturer basic data on ledger: %v", err)
	}

	return nil
}

func (cc *TextileChaincode) initApparelManufacturers(ctx contractapi.TransactionContextInterface) error {

	apparelManufacturerTransactions := []TransactionsList{
		{
			TransactionDate:          "2024-03-28",
			TransactionRecipient:     "Retailer 1",
			TransactionQuantity:      200.0,
			TransactionFiberType:    "T-shirt",
		},
		{
			TransactionDate:          "2024-03-30",
			TransactionRecipient:     "Retailer 2",
			TransactionQuantity:      300.0,
			TransactionFiberType:    "T-shirt",
		},
	}

	// Initialize Apparel Manufacturer Data
	apparelManufacturerData := ApparelManufacturerData{
		CompanyName:                "Apparel Manufacturer 1",
		FabricTypeUsed:             "Cotton",
		GarmentTypeProduced:        "T-shirt",
		ProductionVolume:           1000,
		ManufacturingProcess:       []string{"Cutting", "Sewing"},
		MaterialWaste:              0.1,
		SizeRange:                  "S-XL",
		Color:                      []string{"White", "Black", "Blue"},
		StyleDetail:                "Short sleeves, crew neck",
		ApparelMachineSpecification: "Specification 1",
		Inventory:                  500,
        ApparelManufacturingDate:   "24-03-2024",
		ApparelManufacturerTransactions:				apparelManufacturerTransactions,
	}

	// Marshal Apparel Manufacturer Data
	apparelManufacturerDataJSON, err := json.Marshal(apparelManufacturerData)
	if err != nil {
		return fmt.Errorf("failed to marshal apparel manufacturer data: %v", err)
	}

	// Put Apparel Manufacturer Data on ledger
	err = ctx.GetStub().PutState("1", apparelManufacturerDataJSON)
	if err != nil {
		return fmt.Errorf("failed to put apparel manufacturer data on ledger: %v", err)
	}

	// Initialize Apparel Manufacturer Basic Data
	apparelManufacturerBasicData := ApparelManufacturerBasicData{
		ApparelManufacturerID:                  "1",
		Name:                "Apparel Manufacturer 1",
		Address:             "Address 1",
		ApparelManufacturer: []ApparelManufacturerData{apparelManufacturerData},
	}

	// Marshal Apparel Manufacturer Basic Data
	apparelManufacturerBasicDataJSON, err := json.Marshal(apparelManufacturerBasicData)
	if err != nil {
		return fmt.Errorf("failed to marshal apparel manufacturer basic data: %v", err)
	}

	// Put Apparel Manufacturer Basic Data on ledger
	err = ctx.GetStub().PutState(apparelManufacturerBasicData.ApparelManufacturerID, apparelManufacturerBasicDataJSON)
	if err != nil {
		return fmt.Errorf("failed to put apparel manufacturer basic data on ledger: %v", err)
	}

	return nil
}

func (cc *TextileChaincode) initRetailers(ctx contractapi.TransactionContextInterface) error {
	// Initialize Garment Inventory
	garmentInventory := []GarmentInventory{
		{Type: "T-shirt", Color: "White", Brand: "Brand A", Size: "S", Quantity: 100},
		{Type: "T-shirt", Color: "White", Brand: "Brand A", Size: "M", Quantity: 150},
		{Type: "T-shirt", Color: "White", Brand: "Brand A", Size: "L", Quantity: 200},
		{Type: "T-shirt", Color: "Black", Brand: "Brand B", Size: "S", Quantity: 120},
		{Type: "T-shirt", Color: "Black", Brand: "Brand B", Size: "M", Quantity: 180},
		{Type: "T-shirt", Color: "Black", Brand: "Brand B", Size: "L", Quantity: 220},
	}

	// Marshal Garment Inventory
	garmentInventoryJSON, err := json.Marshal(garmentInventory)
	if err != nil {
		return fmt.Errorf("failed to marshal garment inventory data: %v", err)
	}

	// Put Garment Inventory on ledger
	err = ctx.GetStub().PutState("1", garmentInventoryJSON)
	if err != nil {
		return fmt.Errorf("failed to put garment inventory data on ledger: %v", err)
	}

	// Initialize Retailer Data
	retailerData := RetailerData{
		StoreName:    "Retailer Store",
		StoreAddress: "Retailer Address",
		ProductAssortment: []string{
			"T-shirt - White - Brand A",
			"T-shirt - Black - Brand B",
		},
		Inventory: garmentInventory,
	}

	// Marshal Retailer Data
	retailerDataJSON, err := json.Marshal(retailerData)
	if err != nil {
		return fmt.Errorf("failed to marshal retailer data: %v", err)
	}

	// Put Retailer Data on ledger
	err = ctx.GetStub().PutState("1", retailerDataJSON)
	if err != nil {
		return fmt.Errorf("failed to put retailer data on ledger: %v", err)
	}

	// Initialize Retailer Basic Data
	retailerBasicData := RetailerBasicData{
		RetailerID:           "1",
		Name:         "Retailer 1",
		Address:      "Retailer Address 1",
		RetailerData: []RetailerData{retailerData},
	}

	// Marshal Retailer Basic Data
	retailerBasicDataJSON, err := json.Marshal(retailerBasicData)
	if err != nil {
		return fmt.Errorf("failed to marshal retailer basic data: %v", err)
	}

	// Put Retailer Basic Data on ledger
	err = ctx.GetStub().PutState(retailerBasicData.RetailerID, retailerBasicDataJSON)
	if err != nil {
		return fmt.Errorf("failed to put retailer basic data on ledger: %v", err)
	}

	return nil
}

// CreateBatch creates a new batch with the provided batch ID.
func (cc *TextileChaincode) CreateBatch(ctx contractapi.TransactionContextInterface, batchID string) error {
	// Extract submitting organization information
    clientID, err := ctx.GetClientIdentity().GetMSPID()
    if err != nil {
        return fmt.Errorf("failed to get client MSP ID: %v", err)
    }

    // Check if submitting organization is Org1
    if clientID != "Org1MSP" {
        return fmt.Errorf("unauthorized organization; Org1 required: %s", clientID)
    }

	// Check if the batch already exists
	exists, err := cc.BatchExist(ctx, batchID)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("batch with ID %s already exists", batchID)
	}

	batch := BatchHistory{
		BatchID:      batchID,
		CurrentStage: "new batch created",
	}

	// Marshal the batch data
	batchJSON, err := json.Marshal(batch)
	if err != nil {
		return fmt.Errorf("failed to marshal batch data: %v", err)
	}

	// Put the batch data on the ledger
	err = ctx.GetStub().PutState(batchID, batchJSON)
	if err != nil {
		return fmt.Errorf("failed to put batch data on ledger: %v", err)
	}

	return nil
}

// BatchExist checks if a batch with the given ID exists on the ledger.
func (cc *TextileChaincode) BatchExist(ctx contractapi.TransactionContextInterface, batchID string) (bool, error) {
	batchJSON, err := ctx.GetStub().GetState(batchID)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}
	return batchJSON != nil, nil
}

// Fabric producer code started:
// CreateFiberProducer creates a new fiber producer with the given data.
func (cc *TextileChaincode) CreateFiberProducer(ctx contractapi.TransactionContextInterface, id, name, address string) error {
	// Extract submitting organization information
    clientID, err := ctx.GetClientIdentity().GetMSPID()
    if err != nil {
        return fmt.Errorf("failed to get client MSP ID: %v", err)
    }

    // Check if submitting organization is Org1
    if clientID != "Org1MSP" {
        return fmt.Errorf("unauthorized organization; Org1 required: %s", clientID)
    }
	// Check if fiber producer already exists
	exists, err := cc.FiberProducerExist(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("fiber producer with ID %s already exists", id)
	}

	// Create fiber producer object
	fiberProducer := FiberProducerBasicData{
		FiberProducerID:      id,
		Name:    name,
		Address: address,
		FiberProducer: []FiberProducerData{},
	}

	// Marshal fiber producer data
	fiberProducerJSON, err := json.Marshal(fiberProducer)
	if err != nil {
		return fmt.Errorf("failed to marshal fiber producer data: %v", err)
	}

	// Put fiber producer data on ledger
	err = ctx.GetStub().PutState(id, fiberProducerJSON)
	if err != nil {
		return fmt.Errorf("failed to put fiber producer data on ledger: %v", err)
	}

	return nil
}

// FiberProducerExist checks if a fiber producer with the given ID exists on the ledger.
func (cc *TextileChaincode) FiberProducerExist(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	fiberProducerJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}
	return fiberProducerJSON != nil, nil
}

func (cc *TextileChaincode) AddFiberProducerData(ctx contractapi.TransactionContextInterface, fiberProducerID string, batchID string, producerName string, fiberType string, location string, harvestDate string, landArea float64, productionCapacity int, farmingPractices string, soilHealth string, yield int, fertilizerName string, fertilizerQuantity int, pesticideName string, pesticideQuantity int, herbicideName string, herbicideQuantity int, waterUsage float64, inventory float64, fiberProducingDate string, transactions []TransactionsList) error {
    // // Extract submitting organization information
    // clientID, err := ctx.GetClientIdentity().GetMSPID()
    // if err != nil {
    //     return fmt.Errorf("failed to get client MSP ID: %v", err)
    // }

    // // Check if submitting organization is Org1
    // if clientID != "Org1MSP" {
    //     return fmt.Errorf("unauthorized organization; Org1 required: %s", clientID)
    // }

	// // Retrieve batch history from the ledger
	// batchHistoryJSON, err := ctx.GetStub().GetState(batchID)
    // var batchHistory BatchHistory

    // if err != nil {
    //     return fmt.Errorf("failed to read batch history from the ledger: %v", err)
    // }

    // if batchHistoryJSON != nil {
    //     // Batch exists, so unmarshal it
    //     if err := json.Unmarshal(batchHistoryJSON, &batchHistory); err != nil {
    //         return fmt.Errorf("Errored occur while unmarshalling the data: %v", err)
    //     }
    // } else {
    //     // Batch doesn't exist, so create it
    //     if err := cc.CreateBatch(ctx, batchID); err != nil {
    //         return fmt.Errorf("failed to create new batch: %v", err)
    //     }
    //     fmt.Printf("Batch created successfully!!")

    //     // Fetch the batch history data again after creating the batch
    //     batchHistoryJSON, err = ctx.GetStub().GetState(batchID)
    //     if err != nil {
    //         return fmt.Errorf("failed to read batch history from the ledger after creating the batch: %v", err)
    //     }

    //     // Unmarshal the newly fetched batch history data
    //     if batchHistoryJSON != nil {
    //         if err := json.Unmarshal(batchHistoryJSON, &batchHistory); err != nil {
    //             return fmt.Errorf("error occurred while unmarshalling the data after creating the batch: %v", err)
    //         }
    //     } else {
    //         return fmt.Errorf("failed to fetch batch history data after creating the batch")
    //     }
    // }

	// Extract submitting organization information
    clientID, err := ctx.GetClientIdentity().GetMSPID()
    if err != nil {
        return fmt.Errorf("failed to get client MSP ID: %v", err)
    }

    // Check if submitting organization is Org1
    if clientID != "Org1MSP" {
        return fmt.Errorf("unauthorized organization; Org1 required: %s", clientID)
    }

    // Check if the batch already exists
    batchExists, err := cc.BatchExist(ctx, batchID)
    if err != nil {
        return fmt.Errorf("failed to check batch existence: %v", err)
    }

    if !batchExists {
        // Batch doesn't exist, so create it
        if err := cc.CreateBatch(ctx, batchID); err != nil {
            return fmt.Errorf("failed to create new batch: %v", err)
        }
        fmt.Println("Batch created successfully!!")
    } else {
        fmt.Println("Batch already exists.")
    }

    // Retrieve batch history from the ledger
    batchHistoryJSON, err := ctx.GetStub().GetState(batchID)
    if err != nil {
        return fmt.Errorf("failed to read batch history from the ledger: %v", err)
    }

    var batchHistory BatchHistory

    if batchHistoryJSON != nil {
        // Batch exists, so unmarshal it
        if err := json.Unmarshal(batchHistoryJSON, &batchHistory); err != nil {
            return fmt.Errorf("failed to unmarshal batch history data: %v", err)
        }
    } else {
        return fmt.Errorf("batch history data is nil")
    }

    // Update current stage of the batch
    batchHistory.CurrentStage = "Fiber Producer Stage"

    // Create FiberProducerData
    fiberProducerData := FiberProducerData{
        BatchID:             batchID,
        ProducerName:        producerName,
        FiberType:           fiberType,
        Location:            location,
        HarvestDate:         harvestDate,
        LandArea:            landArea,
        ProductionCapacity:  productionCapacity,
        FarmingPractices:    farmingPractices,
        SoilHealth:          soilHealth,
        Yield:               yield,
        FertilizerName:      fertilizerName,
        FertilizerQuantity:  fertilizerQuantity,
        PesticideName:       pesticideName,
        PesticideQuantity:   pesticideQuantity,
        HerbicideName:       herbicideName,
        HerbicideQuantity:   herbicideQuantity,
        WaterUsage:          waterUsage,
        Inventory:           inventory,
        FiberProducingDate:  fiberProducingDate,
        FiberProducerTransactions: transactions,
    }

    // Add FiberProducerData to FiberProducerBatch in BatchHistory
    batchHistory.FiberProducerBatch = fiberProducerData

    // Marshal updated batch history into JSON
    updatedBatchHistoryJSON, err := json.Marshal(batchHistory)
    if err != nil {
        return fmt.Errorf("failed to marshal updated batch history data: %v", err)
    }

    // Update batch history in the ledger
    if err := ctx.GetStub().PutState(batchID, updatedBatchHistoryJSON); err != nil {
        return fmt.Errorf("failed to update batch history in the ledger: %v", err)
    }

    // Retrieve FiberProducerBasicData from the ledger
    fiberProducerBasicDataJSON, err := ctx.GetStub().GetState(fiberProducerID)
    if err != nil {
        return fmt.Errorf("failed to read fiber producer basic data from the ledger: %v", err)
    }
    if fiberProducerBasicDataJSON == nil {
        return fmt.Errorf("fiber producer basic data does not exist")
    }

    // Unmarshal FiberProducerBasicData into FiberProducerBasicData struct
    var fiberProducerBasicData FiberProducerBasicData
    if err := json.Unmarshal(fiberProducerBasicDataJSON, &fiberProducerBasicData); err != nil {
        return fmt.Errorf("failed to unmarshal fiber producer basic data: %v", err)
    }

    // Append FiberProducerData to FiberProducer list in FiberProducerBasicData
    fiberProducerBasicData.FiberProducer = append(fiberProducerBasicData.FiberProducer, fiberProducerData)

    // Marshal updated FiberProducerBasicData into JSON
    updatedFiberProducerBasicDataJSON, err := json.Marshal(fiberProducerBasicData)
    if err != nil {
        return fmt.Errorf("failed to marshal updated fiber producer basic data: %v", err)
    }

    // Update FiberProducerBasicData in the ledger
    if err := ctx.GetStub().PutState(fiberProducerID, updatedFiberProducerBasicDataJSON); err != nil {
        return fmt.Errorf("failed to update fiber producer basic data in the ledger: %v", err)
    }

    return nil
}

func (cc *TextileChaincode) GetFabricProducerFabrics(ctx contractapi.TransactionContextInterface, fiberProducerID string) ([]*FiberProducerData, error) {
    // Extract submitting organization information
    clientID, err := ctx.GetClientIdentity().GetMSPID()
    if err != nil {
        return nil, fmt.Errorf("failed to get client MSP ID: %v", err)
    }

    // Check if submitting organization is Org1
    if clientID != "Org1MSP" {
        return nil, fmt.Errorf("unauthorized organization; Org1 required: %s", clientID)
    }
	// Retrieve fiber producer data from the ledger
    fiberProducerDataJSON, err := ctx.GetStub().GetState(fiberProducerID)
    if err != nil {
        return nil, fmt.Errorf("failed to read fiber producer data from the ledger: %v", err)
    }
    if fiberProducerDataJSON == nil {
        return nil, fmt.Errorf("fiber producer with ID %s does not exist", fiberProducerID)
    }

    // Unmarshal fiber producer data into FiberProducerBasicData struct
    var fiberProducerBasicData FiberProducerBasicData
    if err := json.Unmarshal(fiberProducerDataJSON, &fiberProducerBasicData); err != nil {
        return nil, fmt.Errorf("failed to unmarshal fiber producer data: %v", err)
    }

    // Create a slice to store fabrics produced by the fiber producer
    fabrics := make([]*FiberProducerData, len(fiberProducerBasicData.FiberProducer))

    // Iterate over each fiber producer data to extract fabrics
    for _, producerData := range fiberProducerBasicData.FiberProducer {
        fabrics = append(fabrics, &producerData)
    }

    return fabrics, nil
}
// Fabric producer code endeed.

// Yarn manufacturer code started:
func (cc *TextileChaincode) CreateYarnManufacturer(ctx contractapi.TransactionContextInterface, id, name, address string) error {
	// Extract submitting organization information
    clientID, err := ctx.GetClientIdentity().GetMSPID()
    if err != nil {
        return fmt.Errorf("failed to get client MSP ID: %v", err)
    }

    // Check if submitting organization is Org1
    if clientID != "Org2MSP" {
        return fmt.Errorf("unauthorized organization; Org2 required: %s", clientID)
    }

	// Check if yarn manufacturer already exists
	exists, err := cc.YarnManufacturerExist(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("yarn manufacturer with ID %s already exists", id)
	}

	// Create yarn manufacturer object
	yarnManufacturer := YarnManufacturerBasicData{
		YarnManufacturerID:                id,
		Name:              name,
		Address:           address,
		YarnManufacturer: []YarnManufacturerData{},
	}

	// Marshal yarn manufacturer data
	yarnManufacturerJSON, err := json.Marshal(yarnManufacturer)
	if err != nil {
		return fmt.Errorf("failed to marshal yarn manufacturer data: %v", err)
	}

	// Put yarn manufacturer data on ledger
	err = ctx.GetStub().PutState(id, yarnManufacturerJSON)
	if err != nil {
		return fmt.Errorf("failed to put yarn manufacturer data on ledger: %v", err)
	}

	return nil
}

func (cc *TextileChaincode) YarnManufacturerExist(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	yarnManufacturerJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}
	return yarnManufacturerJSON != nil, nil
}

func (cc *TextileChaincode) AddYarnManufacturerData(ctx contractapi.TransactionContextInterface, yarnManufacturerID string, batchID string, companyName string, fiberTypeUsed string, fiberSource string, yarnTypeProduced string, yarnCount string, productionCapacity int, dyeing string, wasteManagement float64, finishingTreatment string, strength string, colorsUsed []string, twistLevel string, energyConsumption float64, machinery string, inventory float64, yarnManufacturingDate string, transactions []TransactionsList) error {
    // Extract submitting organization information
    clientID, err := ctx.GetClientIdentity().GetMSPID()
    if err != nil {
        return fmt.Errorf("failed to get client MSP ID: %v", err)
    }

    // Check if submitting organization is Org1
    if clientID != "Org2MSP" {
        return fmt.Errorf("unauthorized organization; Org2 required: %s", clientID)
    }

	// Retrieve batch history from the ledger
    batchHistoryJSON, err := ctx.GetStub().GetState(batchID)
    if err != nil {
        return fmt.Errorf("failed to read batch history from the ledger: %v", err)
    }
    if batchHistoryJSON == nil {
        return fmt.Errorf("batch with ID %s does not exist", batchID)
    }

    // Unmarshal batch history into BatchHistory struct
    var batchHistory BatchHistory
    if err := json.Unmarshal(batchHistoryJSON, &batchHistory); err != nil {
        return fmt.Errorf("failed to unmarshal batch history data: %v", err)
    }

    // Update current stage of the batch
    batchHistory.CurrentStage = "Yarn Manufacturer Stage"

    // Create YarnManufacturerData
    yarnManufacturerData := YarnManufacturerData{
        BatchID:                   batchID,
        CompanyName:              companyName,
        FiberTypeUsed:            fiberTypeUsed,
        FiberSource:              fiberSource,
        YarnTypeProduced:         yarnTypeProduced,
        YarnCount:                yarnCount,
        ProductionCapacity:       productionCapacity,
        Dyeing:                   dyeing,
        WasteManagement:          wasteManagement,
        FinishingTreatment:       finishingTreatment,
        Strength:                 strength,
        ColorsUsed:               colorsUsed,
        TwistLevel:               twistLevel,
        EnergyConsumption:        energyConsumption,
        Machinery:                machinery,
        Inventory:                inventory,
        YarnManufacturingDate:    yarnManufacturingDate,
        YarnManufacturerTransactions: transactions,
    }

    // Add YarnManufacturerData to BatchHistory
    batchHistory.YarnManufacturerBatch = yarnManufacturerData

    // Marshal updated batch history into JSON
    updatedBatchHistoryJSON, err := json.Marshal(batchHistory)
    if err != nil {
        return fmt.Errorf("failed to marshal updated batch history data: %v", err)
    }

    // Update batch history in the ledger
    if err := ctx.GetStub().PutState(batchID, updatedBatchHistoryJSON); err != nil {
        return fmt.Errorf("failed to update batch history in the ledger: %v", err)
    }

    // Retrieve YarnManufacturerBasicData from the ledger
    yarnManufacturerBasicDataJSON, err := ctx.GetStub().GetState(yarnManufacturerID)
    if err != nil {
        return fmt.Errorf("failed to read yarn manufacturer basic data from the ledger: %v", err)
    }
    if yarnManufacturerBasicDataJSON == nil {
        return fmt.Errorf("yarn manufacturer basic data does not exist")
    }

    // Unmarshal YarnManufacturerBasicData into YarnManufacturerBasicData struct
    var yarnManufacturerBasicData YarnManufacturerBasicData
    if err := json.Unmarshal(yarnManufacturerBasicDataJSON, &yarnManufacturerBasicData); err != nil {
        return fmt.Errorf("failed to unmarshal yarn manufacturer basic data: %v", err)
    }

    // Append YarnManufacturerData to YarnManufacturer list in YarnManufacturerBasicData
    yarnManufacturerBasicData.YarnManufacturer = append(yarnManufacturerBasicData.YarnManufacturer, yarnManufacturerData)

    // Marshal updated YarnManufacturerBasicData into JSON
    updatedYarnManufacturerBasicDataJSON, err := json.Marshal(yarnManufacturerBasicData)
    if err != nil {
        return fmt.Errorf("failed to marshal updated yarn manufacturer basic data: %v", err)
    }

    // Update YarnManufacturerBasicData in the ledger
    if err := ctx.GetStub().PutState(yarnManufacturerID, updatedYarnManufacturerBasicDataJSON); err != nil {
        return fmt.Errorf("failed to update yarn manufacturer basic data in the ledger: %v", err)
    }

    return nil
}

func (cc *TextileChaincode) GetYarnManufacturerFabrics(ctx contractapi.TransactionContextInterface, yarnManufacturerID string) ([]*YarnManufacturerData, error) {
    // Extract submitting organization information
    clientID, err := ctx.GetClientIdentity().GetMSPID()
    if err != nil {
        return nil, fmt.Errorf("failed to get client MSP ID: %v", err)
    }

    // Check if submitting organization is Org1
    if clientID != "Org2MSP" {
        return nil, fmt.Errorf("unauthorized organization; Org2 required: %s", clientID)
    }

	// Retrieve yarn manufacturer data from the ledger
    yarnManufacturerDataJSON, err := ctx.GetStub().GetState(yarnManufacturerID)
    if err != nil {
        return nil, fmt.Errorf("failed to read yarn manufacturer data from the ledger: %v", err)
    }
    if yarnManufacturerDataJSON == nil {
        return nil, fmt.Errorf("yarn manufacturer with ID %s does not exist", yarnManufacturerID)
    }

    // Unmarshal yarn manufacturer data into YarnManufacturerBasicData struct
    var yarnManufacturerBasicData YarnManufacturerBasicData
    if err := json.Unmarshal(yarnManufacturerDataJSON, &yarnManufacturerBasicData); err != nil {
        return nil, fmt.Errorf("failed to unmarshal yarn manufacturer data: %v", err)
    }

    // Create a slice to store fabrics produced by the yarn manufacturer
    fabrics := make([]*YarnManufacturerData, len(yarnManufacturerBasicData.YarnManufacturer))

    // Iterate over each yarn manufacturer data to extract fabrics
    for _, manufacturerData := range yarnManufacturerBasicData.YarnManufacturer {
        fabrics = append(fabrics, &manufacturerData)
    }

    return fabrics, nil
}
// Yarn manufacturer code ended.

// Fabric manufacturere code started:
func (cc *TextileChaincode) CreateFabricManufacturer(ctx contractapi.TransactionContextInterface, id, name, address string) error {
	// Extract submitting organization information
    clientID, err := ctx.GetClientIdentity().GetMSPID()
    if err != nil {
        return fmt.Errorf("failed to get client MSP ID: %v", err)
    }

    // Check if submitting organization is Org1
    if clientID != "Org3MSP" {
        return fmt.Errorf("unauthorized organization; Org3 required: %s", clientID)
    }
	
	// Check if fabric manufacturer already exists
	exists, err := cc.FabricManufacturerExist(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("fabric manufacturer with ID %s already exists", id)
	}

	// Create fabric manufacturer object
	fabricManufacturer := FabricManufacturerBasicData{
		FabricManufacturerID: id,
		Name:                 name,
		Address:              address,
		FabricManufacturer:   []FabricManufacturerData{},
	}

	// Marshal fabric manufacturer data
	fabricManufacturerJSON, err := json.Marshal(fabricManufacturer)
	if err != nil {
		return fmt.Errorf("failed to marshal fabric manufacturer data: %v", err)
	}

	// Put fabric manufacturer data on ledger
	err = ctx.GetStub().PutState(id, fabricManufacturerJSON)
	if err != nil {
		return fmt.Errorf("failed to put fabric manufacturer data on ledger: %v", err)
	}

	return nil
}

func (cc *TextileChaincode) FabricManufacturerExist(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	yarnManufacturerJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}
	return yarnManufacturerJSON != nil, nil
}

func (cc *TextileChaincode) AddFabricManufacturerData(ctx contractapi.TransactionContextInterface, fabricManufacturerID string, batchID string, companyName string, yarnTypeUsed string, fabricProducedType string, fabricComposition []FiberComposition, finishingTreatment []string, defectRatio float64, fabricColor string, fabricTexture string, machineSpecification string, fabricWaterConsumption float64, productionTechniques []string, fabricManufacturingDate string, inventory float64, transactions []TransactionsList) error {
	// Extract submitting organization information
    clientID, err := ctx.GetClientIdentity().GetMSPID()
    if err != nil {
        return fmt.Errorf("failed to get client MSP ID: %v", err)
    }

    // Check if submitting organization is Org1
    if clientID != "Org3MSP" {
        return fmt.Errorf("unauthorized organization; Org3 required: %s", clientID)
    }
	
	// Retrieve batch history from the ledger
	batchHistoryJSON, err := ctx.GetStub().GetState(batchID)
	if err != nil {
		return fmt.Errorf("failed to read batch history from the ledger: %v", err)
	}
	if batchHistoryJSON == nil {
		return fmt.Errorf("batch with ID %s does not exist", batchID)
	}

	// Unmarshal batch history into BatchHistory struct
	var batchHistory BatchHistory
	if err := json.Unmarshal(batchHistoryJSON, &batchHistory); err != nil {
		return fmt.Errorf("failed to unmarshal batch history data: %v", err)
	}

	// Update current stage of the batch
	batchHistory.CurrentStage = "Fabric Manufacturer Stage"

	// Create FabricManufacturerData
	fabricManufacturerData := FabricManufacturerData{
		BatchID:                   batchID,
		CompanyName:              companyName,
		YarnTypeUsed:             yarnTypeUsed,
		FabricProducedType:       fabricProducedType,
		FabricComposition:        fabricComposition,
		FinishingTreatment:       finishingTreatment,
		DefectRatio:              defectRatio,
		FabricColor:              fabricColor,
		FabricTexture:            fabricTexture,
		MachineSpecification:     machineSpecification,
		FabricWaterConsumption:   fabricWaterConsumption,
		ProductionTechniques:     productionTechniques,
		FabricManufacturingDate:  fabricManufacturingDate,
		Inventory:                inventory,
		FabricProducerTransactions: transactions,
	}

	// Add FabricManufacturerData to BatchHistory
	batchHistory.FabricManufacturerBatch = fabricManufacturerData

	// Marshal updated batch history into JSON
	updatedBatchHistoryJSON, err := json.Marshal(batchHistory)
	if err != nil {
		return fmt.Errorf("failed to marshal updated batch history data: %v", err)
	}

	// Update batch history in the ledger
	if err := ctx.GetStub().PutState(batchID, updatedBatchHistoryJSON); err != nil {
		return fmt.Errorf("failed to update batch history in the ledger: %v", err)
	}

	// Retrieve FabricManufacturerBasicData from the ledger
	fabricManufacturerBasicDataJSON, err := ctx.GetStub().GetState(fabricManufacturerID)
	if err != nil {
		return fmt.Errorf("failed to read fabric manufacturer basic data from the ledger: %v", err)
	}
	if fabricManufacturerBasicDataJSON == nil {
		return fmt.Errorf("fabric manufacturer basic data does not exist")
	}

	// Unmarshal FabricManufacturerBasicData into FabricManufacturerBasicData struct
	var fabricManufacturerBasicData FabricManufacturerBasicData
	if err := json.Unmarshal(fabricManufacturerBasicDataJSON, &fabricManufacturerBasicData); err != nil {
		return fmt.Errorf("failed to unmarshal fabric manufacturer basic data: %v", err)
	}

	// Append FabricManufacturerData to FabricManufacturer list in FabricManufacturerBasicData
	fabricManufacturerBasicData.FabricManufacturer = append(fabricManufacturerBasicData.FabricManufacturer, fabricManufacturerData)

	// Marshal updated FabricManufacturerBasicData into JSON
	updatedFabricManufacturerBasicDataJSON, err := json.Marshal(fabricManufacturerBasicData)
	if err != nil {
		return fmt.Errorf("failed to marshal updated fabric manufacturer basic data: %v", err)
	}

	// Update FabricManufacturerBasicData in the ledger
	if err := ctx.GetStub().PutState(fabricManufacturerID, updatedFabricManufacturerBasicDataJSON); err != nil {
		return fmt.Errorf("failed to update fabric manufacturer basic data in the ledger: %v", err)
	}

	return nil
}

func (cc *TextileChaincode) GetFabricManufacturerFabrics(ctx contractapi.TransactionContextInterface, fabricManufacturerID string) ([]*FabricManufacturerData, error) {
    // Extract submitting organization information
    clientID, err := ctx.GetClientIdentity().GetMSPID()
    if err != nil {
        return nil, fmt.Errorf("failed to get client MSP ID: %v", err)
    }

    // Check if submitting organization is Org1
    if clientID != "Org3MSP" {
        return nil, fmt.Errorf("unauthorized organization; Org3 required: %s", clientID)
    }
	
	// Retrieve fabric manufacturer data from the ledger
    fabricManufacturerDataJSON, err := ctx.GetStub().GetState(fabricManufacturerID)
    if err != nil {
        return nil, fmt.Errorf("failed to read fabric manufacturer data from the ledger: %v", err)
    }
    if fabricManufacturerDataJSON == nil {
        return nil, fmt.Errorf("fabric manufacturer with ID %s does not exist", fabricManufacturerID)
    }

    // Unmarshal fabric manufacturer data into FabricManufacturerBasicData struct
    var fabricManufacturerBasicData FabricManufacturerBasicData
    if err := json.Unmarshal(fabricManufacturerDataJSON, &fabricManufacturerBasicData); err != nil {
        return nil, fmt.Errorf("failed to unmarshal fabric manufacturer data: %v", err)
    }

    // Create a slice to store fabrics produced by the fabric manufacturer
    fabrics := make([]*FabricManufacturerData, len(fabricManufacturerBasicData.FabricManufacturer))

    // Iterate over each fabric manufacturer data to extract fabrics
    for _, manufacturerData := range fabricManufacturerBasicData.FabricManufacturer {
        fabrics = append(fabrics, &manufacturerData)
    }

    return fabrics, nil
}
// Fabric manufacturer code ended.

// Apparel Manufacturer code started:
func (cc *TextileChaincode) CreateApparelManufacturer(ctx contractapi.TransactionContextInterface, id, name, address string) error {
	// Extract submitting organization information
    clientID, err := ctx.GetClientIdentity().GetMSPID()
    if err != nil {
        return fmt.Errorf("failed to get client MSP ID: %v", err)
    }

    // Check if submitting organization is Org1
    if clientID != "Org4MSP" {
        return fmt.Errorf("unauthorized organization; Org4 required: %s", clientID)
    }
	
	// Check if apparel manufacturer already exists
	exists, err := cc.ApparelManufacturerExist(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("apparel manufacturer with ID %s already exists", id)
	}

	// Create apparel manufacturer object
	apparelManufacturer := ApparelManufacturerBasicData{
		ApparelManufacturerID: id,
		Name:                  name,
		Address:               address,
		ApparelManufacturer:   []ApparelManufacturerData{},
	}

	// Marshal apparel manufacturer data
	apparelManufacturerJSON, err := json.Marshal(apparelManufacturer)
	if err != nil {
		return fmt.Errorf("failed to marshal apparel manufacturer data: %v", err)
	}

	// Put apparel manufacturer data on ledger
	err = ctx.GetStub().PutState(id, apparelManufacturerJSON)
	if err != nil {
		return fmt.Errorf("failed to put apparel manufacturer data on ledger: %v", err)
	}

	return nil
}

func (cc *TextileChaincode) ApparelManufacturerExist(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	apparelManufacturerJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}
	return apparelManufacturerJSON != nil, nil
}

func (cc *TextileChaincode) AddApparelManufacturerData(ctx contractapi.TransactionContextInterface, apparelManufacturerID string, batchID string, companyName string, fabricTypeUsed string, garmentTypeProduced string, productionVolume int, manufacturingProcess []string, materialWaste float64, sizeRange string, color []string, styleDetail string, apparelMachineSpecification string, inventory float64, apparelManufacturingDate string, transactions []TransactionsList) error {
	// Extract submitting organization information
    clientID, err := ctx.GetClientIdentity().GetMSPID()
    if err != nil {
        return fmt.Errorf("failed to get client MSP ID: %v", err)
    }

    // Check if submitting organization is Org1
    if clientID != "Org4MSP" {
        return fmt.Errorf("unauthorized organization; Org4 required: %s", clientID)
    }
	
	// Retrieve batch history from the ledger
	batchHistoryJSON, err := ctx.GetStub().GetState(batchID)
	if err != nil {
		return fmt.Errorf("failed to read batch history from the ledger: %v", err)
	}
	if batchHistoryJSON == nil {
		return fmt.Errorf("batch with ID %s does not exist", batchID)
	}

	// Unmarshal batch history into BatchHistory struct
	var batchHistory BatchHistory
	if err := json.Unmarshal(batchHistoryJSON, &batchHistory); err != nil {
		return fmt.Errorf("failed to unmarshal batch history data: %v", err)
	}

	// Update current stage of the batch
	batchHistory.CurrentStage = "Apparel Manufacturer Stage"

	// Create ApparelManufacturerData
	apparelManufacturerData := ApparelManufacturerData{
		BatchID:                        batchID,
		CompanyName:                   companyName,
		FabricTypeUsed:                fabricTypeUsed,
		GarmentTypeProduced:           garmentTypeProduced,
		ProductionVolume:              productionVolume,
		ManufacturingProcess:          manufacturingProcess,
		MaterialWaste:                 materialWaste,
		SizeRange:                     sizeRange,
		Color:                         color,
		StyleDetail:                   styleDetail,
		ApparelMachineSpecification:   apparelMachineSpecification,
		Inventory:                     inventory,
		ApparelManufacturingDate:      apparelManufacturingDate,
		ApparelManufacturerTransactions: transactions,
	}

	// Add ApparelManufacturerData to BatchHistory
	batchHistory.ApparelManufacturerBatch = apparelManufacturerData

	// Marshal updated batch history into JSON
	updatedBatchHistoryJSON, err := json.Marshal(batchHistory)
	if err != nil {
		return fmt.Errorf("failed to marshal updated batch history data: %v", err)
	}

	// Update batch history in the ledger
	if err := ctx.GetStub().PutState(batchID, updatedBatchHistoryJSON); err != nil {
		return fmt.Errorf("failed to update batch history in the ledger: %v", err)
	}

	// Retrieve ApparelManufacturerBasicData from the ledger
	apparelManufacturerBasicDataJSON, err := ctx.GetStub().GetState(apparelManufacturerID)
	if err != nil {
		return fmt.Errorf("failed to read apparel manufacturer basic data from the ledger: %v", err)
	}
	if apparelManufacturerBasicDataJSON == nil {
		return fmt.Errorf("apparel manufacturer basic data does not exist")
	}

	// Unmarshal ApparelManufacturerBasicData into ApparelManufacturerBasicData struct
	var apparelManufacturerBasicData ApparelManufacturerBasicData
	if err := json.Unmarshal(apparelManufacturerBasicDataJSON, &apparelManufacturerBasicData); err != nil {
		return fmt.Errorf("failed to unmarshal apparel manufacturer basic data: %v", err)
	}

	// Append ApparelManufacturerData to ApparelManufacturer list in ApparelManufacturerBasicData
	apparelManufacturerBasicData.ApparelManufacturer = append(apparelManufacturerBasicData.ApparelManufacturer, apparelManufacturerData)

	// Marshal updated ApparelManufacturerBasicData into JSON
	updatedApparelManufacturerBasicDataJSON, err := json.Marshal(apparelManufacturerBasicData)
	if err != nil {
		return fmt.Errorf("failed to marshal updated apparel manufacturer basic data: %v", err)
	}

	// Update ApparelManufacturerBasicData in the ledger
	if err := ctx.GetStub().PutState(apparelManufacturerID, updatedApparelManufacturerBasicDataJSON); err != nil {
		return fmt.Errorf("failed to update apparel manufacturer basic data in the ledger: %v", err)
	}

	return nil
}

func (cc *TextileChaincode) GetApparelManufacturerFabrics(ctx contractapi.TransactionContextInterface, apparelManufacturerID string) ([]*ApparelManufacturerData, error) {
    // Extract submitting organization information
    clientID, err := ctx.GetClientIdentity().GetMSPID()
    if err != nil {
        return nil, fmt.Errorf("failed to get client MSP ID: %v", err)
    }

    // Check if submitting organization is Org1
    if clientID != "Org4MSP" {
        return nil, fmt.Errorf("unauthorized organization; Org4 required: %s", clientID)
    }
	
	// Retrieve apparel manufacturer data from the ledger
    apparelManufacturerDataJSON, err := ctx.GetStub().GetState(apparelManufacturerID)
    if err != nil {
        return nil, fmt.Errorf("failed to read apparel manufacturer data from the ledger: %v", err)
    }
    if apparelManufacturerDataJSON == nil {
        return nil, fmt.Errorf("apparel manufacturer with ID %s does not exist", apparelManufacturerID)
    }

    // Unmarshal apparel manufacturer data into ApparelManufacturerBasicData struct
    var apparelManufacturerBasicData ApparelManufacturerBasicData
    if err := json.Unmarshal(apparelManufacturerDataJSON, &apparelManufacturerBasicData); err != nil {
        return nil, fmt.Errorf("failed to unmarshal apparel manufacturer data: %v", err)
    }

    // Create a slice to store fabrics produced by the apparel manufacturer
    fabrics := make([]*ApparelManufacturerData, len(apparelManufacturerBasicData.ApparelManufacturer))

    // Iterate over each apparel manufacturer data to extract fabrics
    for _, manufacturerData := range apparelManufacturerBasicData.ApparelManufacturer {
        fabrics = append(fabrics, &manufacturerData)
    }

    return fabrics, nil
}
// Apparel manufacturer code ended.

// Retailer code started:
func (cc *TextileChaincode) CreateRetailer(ctx contractapi.TransactionContextInterface, id, name, address string) error {
	// Extract submitting organization information
    clientID, err := ctx.GetClientIdentity().GetMSPID()
    if err != nil {
        return fmt.Errorf("failed to get client MSP ID: %v", err)
    }

    // Check if submitting organization is Org1
    if clientID != "Org5MSP" {
        return fmt.Errorf("unauthorized organization; Org5 required: %s", clientID)
    }
	
	// Check if retailer already exists
	exists, err := cc.RetailerExist(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("retailer with ID %s already exists", id)
	}

	// Create retailer object
	retailer := RetailerBasicData{
		RetailerID:   id,
		Name:         name,
		Address:      address,
		RetailerData: []RetailerData{},
	}

	// Marshal retailer data
	retailerJSON, err := json.Marshal(retailer)
	if err != nil {
		return fmt.Errorf("failed to marshal retailer data: %v", err)
	}

	// Put retailer data on ledger
	err = ctx.GetStub().PutState(id, retailerJSON)
	if err != nil {
		return fmt.Errorf("failed to put retailer data on ledger: %v", err)
	}

	return nil
}

func (cc *TextileChaincode) RetailerExist(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	retailerJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}
	return retailerJSON != nil, nil
}

func (cc *TextileChaincode) AddRetailerData(ctx contractapi.TransactionContextInterface, retailerID, batchID, storeName, storeAddress string, productAssortment []string, inventory []GarmentInventory) error {
	// Extract submitting organization information
    clientID, err := ctx.GetClientIdentity().GetMSPID()
    if err != nil {
        return fmt.Errorf("failed to get client MSP ID: %v", err)
    }

    // Check if submitting organization is Org1
    if clientID != "Org5MSP" {
        return fmt.Errorf("unauthorized organization; Org5 required: %s", clientID)
    }
	
	// Retrieve batch history from the ledger
	batchHistoryJSON, err := ctx.GetStub().GetState(batchID)
	if err != nil {
		return fmt.Errorf("failed to read batch history from the ledger: %v", err)
	}
	if batchHistoryJSON == nil {
		return fmt.Errorf("batch with ID %s does not exist", batchID)
	}

	// Unmarshal batch history into BatchHistory struct
	var batchHistory BatchHistory
	if err := json.Unmarshal(batchHistoryJSON, &batchHistory); err != nil {
		return fmt.Errorf("failed to unmarshal batch history data: %v", err)
	}

	// Update current stage of the batch
	batchHistory.CurrentStage = "Retailer Stage"

	// Create RetailerData
	retailerData := RetailerData{
		BatchID:           batchID,
		StoreName:         storeName,
		StoreAddress:      storeAddress,
		ProductAssortment: productAssortment,
		Inventory:         inventory,
	}

	// Assign RetailerData to RetailerBatch in BatchHistory
	batchHistory.RetailerBatch = retailerData

	// Marshal updated batch history into JSON
	updatedBatchHistoryJSON, err := json.Marshal(batchHistory)
	if err != nil {
		return fmt.Errorf("failed to marshal updated batch history data: %v", err)
	}

	// Update batch history in the ledger
	if err := ctx.GetStub().PutState(batchID, updatedBatchHistoryJSON); err != nil {
		return fmt.Errorf("failed to update batch history in the ledger: %v", err)
	}

	// Retrieve RetailerBasicData from the ledger
	retailerBasicDataJSON, err := ctx.GetStub().GetState(retailerID)
	if err != nil {
		return fmt.Errorf("failed to read retailer basic data from the ledger: %v", err)
	}
	if retailerBasicDataJSON == nil {
		return fmt.Errorf("retailer basic data does not exist")
	}

	// Unmarshal RetailerBasicData into RetailerBasicData struct
	var retailerBasicData RetailerBasicData
	if err := json.Unmarshal(retailerBasicDataJSON, &retailerBasicData); err != nil {
		return fmt.Errorf("failed to unmarshal retailer basic data: %v", err)
	}

	// Append RetailerData to RetailerData list in RetailerBasicData
	retailerBasicData.RetailerData = append(retailerBasicData.RetailerData, retailerData)

	// Marshal updated RetailerBasicData into JSON
	updatedRetailerBasicDataJSON, err := json.Marshal(retailerBasicData)
	if err != nil {
		return fmt.Errorf("failed to marshal updated retailer basic data: %v", err)
	}

	// Update RetailerBasicData in the ledger
	if err := ctx.GetStub().PutState(retailerID, updatedRetailerBasicDataJSON); err != nil {
		return fmt.Errorf("failed to update retailer basic data in the ledger: %v", err)
	}

	return nil
}

func (cc *TextileChaincode) GetRetailerData(ctx contractapi.TransactionContextInterface, retailerID string) ([]*RetailerData, error) {
    // Extract submitting organization information
    clientID, err := ctx.GetClientIdentity().GetMSPID()
    if err != nil {
        return nil, fmt.Errorf("failed to get client MSP ID: %v", err)
    }

    // Check if submitting organization is Org1
    if clientID != "Org5MSP" {
        return nil, fmt.Errorf("unauthorized organization; Org5 required: %s", clientID)
    }
	
	// Retrieve retailer data from the ledger
    retailerDataJSON, err := ctx.GetStub().GetState(retailerID)
    if err != nil {
        return nil, fmt.Errorf("failed to read retailer data from the ledger: %v", err)
    }
    if retailerDataJSON == nil {
        return nil, fmt.Errorf("retailer with ID %s does not exist", retailerID)
    }

    // Unmarshal retailer data into RetailerBasicData struct
    var retailerBasicData RetailerBasicData
    if err := json.Unmarshal(retailerDataJSON, &retailerBasicData); err != nil {
        return nil, fmt.Errorf("failed to unmarshal retailer data: %v", err)
    }

    // Create a slice to store retailer data
    retailerData := make([]*RetailerData, len(retailerBasicData.RetailerData))

    // Iterate over each retailer data to extract retailer data
    for _, data := range retailerBasicData.RetailerData {
        retailerData = append(retailerData, &data)
    }

    return retailerData, nil
}
// Retailer code ended.

// This function will retrieve the data for a batch
func (cc *TextileChaincode) GetBatchData(ctx contractapi.TransactionContextInterface, batchID string) (*BatchHistory, error) {

    batchDataJSON, err := ctx.GetStub().GetState(batchID)
    if err != nil {
        return nil, fmt.Errorf("Failed to read from the ledger: %v", err)
    }

    if batchDataJSON == nil {
        return nil, fmt.Errorf("Batch with ID %s does not exist", batchID)
    }

    var batchData BatchHistory
    if err := json.Unmarshal(batchDataJSON, &batchData); err != nil {
        return nil, fmt.Errorf("Failed to unmarshal BatchHistory: %v", err)
    }

    return &batchData, nil
}

func main() {
	// Create a new Smart Contract
	smartContract, err := contractapi.NewChaincode(&TextileChaincode{})
	if err != nil {
		fmt.Printf("Error creating SmartContract chaincode: %v\n", err)
		return
	}

	if err := smartContract.Start(); err != nil {
		fmt.Printf("Error starting SmartContract chaincode: %v\n", err)
		return
	}
}

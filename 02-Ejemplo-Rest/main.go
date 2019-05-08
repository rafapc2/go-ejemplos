package main
import (
  "fmt"
  "encoding/json"
  "log"
  "net/http"
  "github.com/gorilla/mux"
)

type Customer struct{
  Id string `json:"id,omitempty"`
  FirstName string `json:"firstName,omitempty"`
  LastName string `json:"lastName,omitempty"`
  Gender string `json:"gender,omitempty"`
  Addres *Addres `json:"addres,omitempty"`
}

type Addres struct{

  Country string `json:"country,omitempty"`
  City string `json:"city,omitempty"`
  State string `json:"state,omitempty"`
  Street string `json:"street,omitempty"`
  StreetNumber string `json:"streetNumber,omitempty"`
  PostalCode int `json:"postalCode,omitempty"`

}

var customerList []Customer

func getCustomersEndPoint(resp http.ResponseWriter, req *http.Request){
  fmt.Println("procesando getCustomersEndPoint")
  json.NewEncoder(resp).Encode(customerList)

}

func getCustomerEndPoint(resp http.ResponseWriter, req *http.Request){
  params := mux.Vars(req)
  for _, item := range customerList {
    if item.Id == params["id"] {
      json.NewEncoder(resp).Encode(item)
      return
    }
  }
  //si el id no existe se retorna vacio
  json.NewEncoder(resp).Encode(&Customer{})
}

func createCustomersEndPoint(resp http.ResponseWriter, req *http.Request){
  params := mux.Vars(req)
  var newCust Customer
  _= json.NewDecoder(req.Body).Decode(&newCust)

  newCust.Id = params["id"]
  customerList = append(customerList, newCust)
  json.NewEncoder(resp).Encode(customerList)


}
func deleteCustomersEndPoint(resp http.ResponseWriter, req *http.Request){

}

/* version 1.0
primer ejemplo de un servicio rest/http */
func main(){
  log.Println("Inicando Server")
  //initialize(&customerList)
  router := mux.NewRouter()

  customerList = append(customerList,Customer{Id:"001",FirstName:"Joaco",LastName:"Prudencio",Gender:"M",
    Addres: &Addres{Country:"Chile",City:"Santiago",State:"RM",Street:"Teresa Vial",StreetNumber:"1431",PostalCode:120923}})
  customerList = append(customerList,Customer{Id:"002",FirstName:"Cristobal",LastName:"Henry",Gender:"M",
    Addres: &Addres{Country:"Chile",City:"Santiago",State:"RM",Street:"B Buston",StreetNumber:"986",PostalCode:387245}})
  customerList = append(customerList,Customer{Id:"003",FirstName:"Maria",LastName:"Perez",Gender:"F",
    Addres: &Addres{Country:"Chile",City:"Santiago",State:"RM",Street:"Alameda",StreetNumber:"2122",PostalCode:334421}})

  //endpoints
  router.HandleFunc("/customers", getCustomersEndPoint).Methods("GET")
  router.HandleFunc("/customers/{id}", getCustomerEndPoint).Methods("GET")
  router.HandleFunc("/customers/{id}", createCustomersEndPoint).Methods("POST")
  router.HandleFunc("/customers/{id}", deleteCustomersEndPoint).Methods("DELETE")

  log.Fatal(http.ListenAndServe(":8080", router))

}

func initialize(customerList *[]Customer ){


}

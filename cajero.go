package main

import "fmt"
import "io/ioutil"
import "strings"
import "strconv"
import "os"
import "os/exec"

const SIZE_ROW_CLIENTE = 10 //limite de clientes (clientes.txt)
const SIZE_COLUMN_CLIENTE = 4 //limite de columnas (clientes.txt)

func main(){
	var operacion int
	operacion = openInit()
	initOperacion(operacion)
}

func openInit() int{
	var operacion int 
	fmt.Println("Bienvenido a la banca Go")
	fmt.Println("Menu, Seleccione el tipo de operacion:")
	fmt.Println("1) Cliente, retiros y transferencias")
	fmt.Println("2) Gestor banca")
	fmt.Scanf("%d\n",&operacion)
	return operacion
}

func initOperacion(operacion int){
	if (operacion == 1) {
		var cuenta int 
		var nip int
		fmt.Println("Ingresar su numero de cuenta")
		fmt.Scanf("%d\n",&cuenta)
		fmt.Println("Ingresar su nip")
		fmt.Scanf("%d\n",&nip)
		//validar si es cliente
		//-----------------------------
		var credenciales bool
		credenciales = validarCliente(cuenta,nip)
		if(credenciales == true){
			fmt.Println("Acceso correcto")
		}else{
			cleanScreen()
			fmt.Println("Acceso denegado, intente de nuevo")
			main()
		}
		//-----------------------------
	}else if (operacion == 2){
		var clave int
		fmt.Println("Ingresar la clave general")
		fmt.Scanf("%d\n",&clave)
		//validar el acceso
		fmt.Println("msg: 2")
	}else{
		cleanScreen()
		fmt.Println("La opcion seleccionada es incorrecta")
		main()
	}
}

func validarCliente(cuenta int,nip int) bool{
	var estatus bool
	estatus = false
	infocliente, err := ioutil.ReadFile("files/clientes.txt")
	showError(err)
	var info string
	info = string(infocliente)
	var cantidadFilas int
	cantidadFilas = len(strings.Split(info,"\n"))
	var fila[] string
	fila = strings.Split(info,"\n")
	if fila == nil{
		fmt.Println("Filas vacias")
	}else{
		var columna [] string
		var arreglo [SIZE_ROW_CLIENTE][SIZE_COLUMN_CLIENTE] string
		for i := 0; i < cantidadFilas; i++{
			columna = strings.Split(fila[i],",")
			for j:= 0;j<=3;j++{
				arreglo[i][j] = columna[j]
			}
		}
		if columna == nil{
			fmt.Println("Columnas vacias")
		}else{
			for i := 0; i < cantidadFilas; i++{
				if((strconv.FormatInt(int64(cuenta),10) == arreglo[i][0]) && (strconv.FormatInt(int64(nip),10) == arreglo[i][2])) {
					estatus = true
				}
			}
		}
	}
	return estatus
}

func showError(e error){
	if(e != nil){
		panic(e)
	}
}

func cleanScreen(){
	clean := exec.Command("clear")
	clean.Stdout = os.Stdout
	clean.Run()
}
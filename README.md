# Prueba Cajero
==========
## Elaborar un programa que permita simular el funcionamiento de un cajero automático, el programa deberá leer 3 archivos de texto, los cuales son los siguientes:
--------------------
> clientes.txt (cuenta,cliente,nip,saldo)
1,luis,1234,7000
2,maria,1234,5000
3,carolina,0000,1000

> cash.txt (denominacion, cantidad)
50,10
100,10
200,10
500,10
1000,10

> config.txt (atributo, valor)
clave,1616
limite,7000 (limite de retiro)


* El programa deberá validar el cliente, solicitando el numero de cuenta y el nip para realizar las transacciones (Retiro).

* El programa deberá validar el cliente, solicitando el numero de cuenta y el nip para realizar transacciones de transferencia a otro cliente solicitando el numero de cuenta del receptor y la cantidad
solicitada.

* El programa permitirá agregar cash al cajero, solicitando la clave general del cajero, denominación del
billete y cantidad.

* El programa permitirá agregar clientes solicitando la clave general del cajero.

* Tu puedes mejorar y agregar mas funcionalidad al cajero, la idea es optimizarlo y agregarle cosas.
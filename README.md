- Hacer un servicio HTTP que exponga un recurso llamado /ping, responda con un 200 OK y el texto "pong" cuando reciba un GET

  	golang caracteristicas (typed, no oo but interfaces, compilado, no virtual machine, runtime, GC, crosscompile)
	Workspace, estructura de directorios, paquetes
	enfasis en parallel and concurrent
	testing framework
	main function
	modificadores de acceso

- Exponer un recurso /category/handlingtime que responda con 400 error un POST

- Exponer un recurso /category/handlingtime que responda con 200 cuando reciba un order_id as json
	
	golang objects (Hacer un obj que reciba los requests.)
	goroutines, channels, mutex, defer

- Exponer un recurso /category/handlingtime que al recibir un POST con order_id, haga un PUT a /internal/seller_reputation/orders/${orderId}/handling_time?caller.scopes=admin
	use composition
	interfaces


-----///DEBUGGER!!!


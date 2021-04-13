
function printHello() {
	while (true) {
		console.log("hello world")
	}	
}

async function main () {
	await printHello()
	setTimeout(() => {
		console.log("the end")
	}, 2000)
}

// will print hello forever. Won't be able to do anything else
// comnpare with not_blocking.go
main()

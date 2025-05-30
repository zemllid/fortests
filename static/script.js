document.addEventListener("DOMContentLoaded", async () => {await loadFrogs();})

async function createMessage() {
    var inputID = document.getElementById("inputName").value;
    var inputa = document.getElementById("inputSpecies").value;
    var inputb = document.getElementById("inputHabitat").value;
    var inputc = document.getElementById("inputAge").value;
    console.log(inputID + inputa + inputb + inputc)
    let response = await fetch('/frogs',{
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify({ name: inputID, species: inputa, habitat: inputb, age: parseInt(inputc) })
      });
}

async function loadFrogs() {
    let response = await fetch('/frogs');  
    if (response.ok){
        let json = await response.json()
        let element = document.getElementById("frogsTableBody");
		
        let plain = document.getElementById("plain");
        element.innerHtml = ""
		plain.innerHtml = ""
		if (json==null) {
			plain.innerHTML= "Жаб нет :("
		} else {
        console.log(json.length)
		let a=""
		let b=""
        for(var i=0;i<json.length; i++) {
            b+= "<td> <tr>" + json[i].id + "<tr>"
			+ json[i].name + "<tr>" 
                                                         + json[i].species + "<tr>" 
                                                         + json[i].habitat + "<tr>"  
                                                         + json[i].age
			a += "\n <br>" + json[i].id + " "
			+ json[i].name + "  " 
                                                         + json[i].species + "  " 
                                                         + json[i].habitat + "  "  
                                                         + json[i].age											 
            console.log(json[i].name + "  " +  json[i].species + "  " + json[i].habitat)
        };
		plain.innerHTML=a
		element.innerHtml = b
		}
    } else{
        alert("Error: " + response.json())
    }  
}

async function delMessage() {
    var inputID = document.getElementById("inputName").value;
    let response = await fetch('/frogs/'+inputID,{
        method: 'DELETE'
      });
}
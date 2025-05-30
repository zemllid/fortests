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

async function getMessage() {
    let response = await fetch('/frogs');  
    if (response.ok){
        let json = await response.json()
        let element = document.getElementById("table");
        element.innerText = ""
        console.log(json.length)
        for(var i=0;i<json.length; i++) {
            element.innerText = element.innerText + "\n" + json[i].name + "  " 
                                                         + json[i].species + "  " 
                                                         + json[i].habitat + "  "  
                                                         + json[i].age
            console.log(json[i].name + "  " +  json[i].species + "  " + json[i].habitat)
        };

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
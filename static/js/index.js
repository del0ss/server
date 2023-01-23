// const btn = document.querySelector(".form > button")
// if (btn) {
//     btn.onclick = e => {
//         const inputs = document.querySelectorAll(".form > input")
//         let data = {}
//         for (let i = 0; i < inputs.length; i++) {
//             data[inputs[i].name] = inputs[i].value
//         }
//
//         const xhr = new XMLHttpRequest()
//         xhr.open("POST", "/auth/reg")
//         xhr.onload = e => {
//            // const response = JSON.parse(e.target.response)
//            // if ("Error" in response) {
//            //      if (response.Error == null) {
//            //          console.log("Пользователь зареган")
//            //      } else {
//            //          console.log(response.Error)
//            //      }
//            // }
//         }
//         xhr.send(JSON.stringify(data))
//         setTimeout(() => document.location.href= "http://localhost:8000/", 100)
//     }
// }

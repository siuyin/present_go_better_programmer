
function A() {
  setInterval(function(){ // cannot have blocking sleep in js
    console.log("A");
  },1000)
}

function b() {
  setInterval(function(){
    console.log("b");
  },1000)
}

A();
b();

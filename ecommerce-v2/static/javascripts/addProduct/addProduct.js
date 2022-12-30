const specification = document.querySelector(".specification");
const addButton = document.querySelector("#addButton");
const specificationField = document.querySelector(".specification-fields");

function removeInput(event) {
	var target = event.target;
	console.log(event.target, "onclick");
	var parent = target.parentElement; //parent of "target"
	parent.remove();
}

function addSpecification() {
	const key = document.createElement("input");
	key.type = "text";
	key.name = "key[]";
	key.placeholder = "key";

	const value = document.createElement("input");
	value.type = "text";
	value.name = "value[]";
	value.placeholder = "value";

	const removeBtn = document.createElement("a");
	removeBtn.className = "delete";
	removeBtn.innerHTML = "&times";

	// removeBtn.addEventListener("click", removeInput);

	const flex = document.createElement("div");
	flex.className = "flex";

	specificationField.appendChild(flex);
	flex.appendChild(key);
	flex.appendChild(value);
	flex.appendChild(removeBtn);
}

specificationField.addEventListener("click", function (e) {
	e.preventDefault();
	if (e.target.classList.contains("delete")) {
		e.target.parentElement.remove();
		console.log(e.target, "removed");
		// document.querySelector(id).scrollIntoView({ behavior: "smooth" });
	}
});

addButton.addEventListener("click", addSpecification);

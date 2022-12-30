const addButton = document.getElementById("addButton");
const removeButton = document.getElementById("");

addButton.addEventListener("click", function () {
	const row = document.querySelector(".row");
	row.insertAdjacentHTML(
		"beforeend",
		`<div class="row">
		<input type="text" name="keys[]" placeholder="key" />
		<input type="text" name="value[]" placeholder="value" />
		<a href="#" id="removeButton">&minus;</a>
	</div>`
	);
	// var specKey = document.createElement("input");
	// specKey.setAttribute("type", "text");
	// specKey.setAttribute("name", "keys[]");
	// specKey.setAttribute("placeholder", "key");

	// var specValue = document.createElement("input");
	// specValue.setAttribute("type", "text");
	// specValue.setAttribute("name", "value[]");
	// specValue.setAttribute("placeholder", "Value");
	// row.appendChild(document.createElement("br"));
	// row.appendChild(specKey);
	// row.appendChild(specValue);
	// console.log("added successfully");
	e.preventDefault();
});

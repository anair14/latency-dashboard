document.addEventListener("DOMContentLoaded", function () {
    const addForm = document.getElementById("add-endpoint-form");
    const removeButtons = document.querySelectorAll(".remove-endpoint");

    if (addForm) {
        addForm.addEventListener("submit", function (event) {
            event.preventDefault();
            const formData = new FormData(addForm);

            fetch("/add-endpoint", {
                method: "POST",
                body: formData,
            })
            .then(response => {
                if (response.ok) {
                    location.reload(); // Refresh page to show new endpoint
                } else {
                    return response.text().then(text => alert("Error: " + text));
                }
            })
            .catch(error => console.error("Error adding endpoint:", error));
        });
    }

    removeButtons.forEach(button => {
        button.addEventListener("click", function () {
            const endpointID = this.dataset.id;

            fetch("/remove-endpoint", {
                method: "POST",
                headers: {
                    "Content-Type": "application/x-www-form-urlencoded",
                },
                body: `endpoint_id=${endpointID}`,
            })
            .then(response => {
                if (response.ok) {
                    location.reload(); // Refresh page after deletion
                } else {
                    return response.text().then(text => alert("Error: " + text));
                }
            })
            .catch(error => console.error("Error removing endpoint:", error));
        });
    });
});


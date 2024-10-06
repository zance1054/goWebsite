let images = [];
let currentIndex = 0;

document.addEventListener("DOMContentLoaded", function () {
    fetch("/api/getGalleryImages")
        .then(response => response.json())
        .then(data => {
            images = data;
            if (images.length > 0) {
                showImage(currentIndex);
            }
        })
        .catch(error => console.error("Error fetching gallery images:", error));
});

function showImage(index) {
    const imgElement = document.getElementById("currentImage");
    imgElement.src = images[index];
    imgElement.style.display = "block"; // Show the image
}

function changeImage(direction) {
    currentIndex += direction;
    
    // Loop around if necessary
    if (currentIndex < 0) {
        currentIndex = images.length - 1; // Go to last image
    } else if (currentIndex >= images.length) {
        currentIndex = 0; // Go back to first image
    }
    
    showImage(currentIndex);
}

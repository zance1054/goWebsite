document.addEventListener("DOMContentLoaded", function () {
    fetch('/blog_posts/blog_posts.json')  // Adjust the path to your JSON file if necessary
    .then(response => response.json())
    .then(posts => {
        const blogContainer = document.getElementById("blog-posts");

        posts.forEach(post => {
            // Create a div for each blog post
            const postDiv = document.createElement("div");
            postDiv.classList.add("fakeimg");

            // Set the background image from the JSON data
            postDiv.style.backgroundImage = `url(${post.backgroundImage})`;
            postDiv.style.backgroundSize = 'cover';
            postDiv.style.backgroundPosition = 'center';

            // Create and add the blog post title
            const title = document.createElement("h2");
            title.innerText = post.title;
            postDiv.appendChild(title);

            // Create and add the blog post date
            const date = document.createElement("h5");
            date.innerText = post.date;
            postDiv.appendChild(date);

            // Append the post div to the blog container
            blogContainer.appendChild(postDiv);
        });
    })
    .catch(error => console.error('Error fetching blog posts:', error));
});

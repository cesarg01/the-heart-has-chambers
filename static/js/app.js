const navSlide = () => {
    const burger = document.querySelector('.burger');
    const nav = document.querySelector('.nav-links');
    const navLinks = document.querySelectorAll('.nav-links li')
    // Get the icon tag
    const icon = burger.querySelector('i');

    // Toggle nav
    burger.addEventListener('click', ()=> {
        nav.classList.toggle('nav-active');
    
        // When the bar icon is clicked change it to times and vice versa
        icon.classList.toggle('fa-times');
        icon.classList.toggle('fa-bars');

         // Animate links
        navLinks.forEach((link, index) => {
            if (link.style.animation) {
                link.style.animation = '';
            }
            else {
                // Adding 1.5 makes a delay when the mobile nav bar opens up.
                link.style.animation = `navLinkFade 0.5s ease forwards ${index / 5 + .5}s`;
            }
        });

    
        
        // Burger animation for the closing icon
        
    });

}

navSlide();
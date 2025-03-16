
const addToFavorites = async (type, id) => {
    try {
        const response = await fetch('/api/favorites', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ type, id }),
        });

        if (response.ok) {
            const button = event.target;
            button.textContent = 'Added to Favorites';
            button.disabled = true;
        }
    } catch (error) {
        console.error('Error adding to favorites:', error);
    }
};

const observeCards = () => {
    const cards = document.querySelectorAll('.content-card');
    const observer = new IntersectionObserver(
        (entries) => {
            entries.forEach(entry => {
                if (entry.isIntersecting) {
                    entry.target.style.opacity = '1';
                    entry.target.style.transform = 'translateY(0)';
                }
            });
        },
        { threshold: 0.1 }
    );

    cards.forEach(card => {
        card.style.opacity = '0';
        card.style.transform = 'translateY(20px)';
        card.style.transition = 'all 0.5s ease-out';
        observer.observe(card);
    });
};

const handleSearch = (event) => {
    event.preventDefault();
    const form = event.target;
    const searchQuery = new FormData(form).get('name');
    const currentUrl = new URL(window.location.href);
    
    currentUrl.searchParams.set('name', searchQuery);
    currentUrl.searchParams.delete('page');
    
    window.location.href = currentUrl.toString();
};

const handleFilter = (filter, value) => {
    const currentUrl = new URL(window.location.href);
    
    if (value) {
        currentUrl.searchParams.set(filter, value);
    } else {
        currentUrl.searchParams.delete(filter);
    }
    
    currentUrl.searchParams.delete('page');
    window.location.href = currentUrl.toString();
};


let lastScroll = 0;
window.addEventListener('scroll', () => {
    const header = document.querySelector('.main-header');
    const currentScroll = window.pageYOffset;

    if (currentScroll > lastScroll) {
        header.style.transform = 'translateY(-100%)';
    } else {
        header.style.transform = 'translateY(0)';
    }

    lastScroll = currentScroll;
});

document.addEventListener('DOMContentLoaded', () => {
    observeCards();
    
    const searchForms = document.querySelectorAll('.search-form');
    searchForms.forEach(form => {
        form.addEventListener('submit', handleSearch);
    });
    
});
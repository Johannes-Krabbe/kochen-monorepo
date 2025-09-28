/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        "./components/**/*.{html,js,templ}",
        "./pages/**/*.{html,js,templ}",
    ],
    theme: {
        extend: {
            fontFamily: {
                sans: ['Lato', 'sans-serif'],
            },
            colors: {
                'grey': 'hsl(var(--color-grey))',
                'light-grey': 'hsl(var(--color-light-grey))',
                'beige': 'hsl(var(--color-beige))',
                'light-brown': 'hsl(var(--color-light-brown))',
                'dark-brown': 'hsl(var(--color-dark-brown))',
                'almost-black': 'hsl(var(--color-almost-black))',
                'background': 'hsl(var(--background))',
                'foreground': 'hsl(var(--foreground))',
                'card': 'hsl(var(--card))',
                'card-foreground': 'hsl(var(--card-foreground))',
                'popover': 'hsl(var(--popover))',
                'popover-foreground': 'hsl(var(--popover-foreground))',
                'primary': 'hsl(var(--primary))',
                'primary-foreground': 'hsl(var(--primary-foreground))',
                'secondary': 'hsl(var(--secondary))',
                'secondary-foreground': 'hsl(var(--secondary-foreground))',
                'muted': 'hsl(var(--muted))',
                'muted-foreground': 'hsl(var(--muted-foreground))',
                'accent': 'hsl(var(--accent))',
                'accent-foreground': 'hsl(var(--accent-foreground))',
                'destructive': 'hsl(var(--destructive))',
                'destructive-foreground': 'hsl(var(--destructive-foreground))',
                'border': 'hsl(var(--border))',
                'input': 'hsl(var(--input))',
                'ring': 'hsl(var(--ring))',
            },
            spacing: {
                'nav': 'var(--nav-bar-height)',
            },
            borderRadius: {
                DEFAULT: 'var(--radius)',
            },
        },
    },
    plugins: [],
}

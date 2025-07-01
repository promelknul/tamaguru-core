module.exports = {
  content: ['./app/**/*.{ts,tsx,js,jsx}'],
  theme: {
    extend: {
      colors: { 'vsx-bg': '#000', 'vsx-primary': '#00ff88' },
      keyframes: { matrix: { '0%,100%': {backgroundPosition:'0 0'}, '50%': {backgroundPosition:'1000px 1000px'} } },
      animation: { matrix: 'matrix 30s linear infinite' }
    }
  }
}

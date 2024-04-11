import { defineStore } from "pinia"

export const useMediaWidth = defineStore('mediaWidth', () => {
  const maxWidth = 1440

  const widgetStyle = { 
    width: `min(${maxWidth}px, 100%)`,
    'margin-left': 'auto',
    'margin-right': 'auto',
    'padding-left': '20px',
    'padding-right': '20px',
  }
  return { maxWidth, widgetStyle }
})
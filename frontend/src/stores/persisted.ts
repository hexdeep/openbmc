import { ref } from 'vue'
import { defineStore } from 'pinia'

export const usePersistedStore = defineStore('persisted', () => {

  const serverAddr = ref('https://axogc.net:8080/api')

  return { serverAddr }
})

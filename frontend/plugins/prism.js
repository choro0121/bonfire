import Prism from 'prismjs/components/prism-core'
import 'prismjs/themes/prism.css'

const languages = require('prismjs/components.js').languages

const loaded = []

function load (lang) {
  if (languages[lang].title !== undefined) {
    if (!loaded.includes(lang)) {
      if (typeof languages[lang].require === 'string') {
        load(languages[lang].require)
      } else if (typeof languages[lang].require === 'object') {
        languages[lang].require.forEach((req) => {
          load(req)
        })
      }
      import('prismjs/components/prism-' + lang)
      loaded.push(lang)
      console.log(lang, languages[lang])
    }
  }
}

Object.keys(languages).forEach((lang) => { load(lang) })

export default (_, inject) => {
  inject('Prism', Prism)
}

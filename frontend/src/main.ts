import './style.css'
import './app.css'

import { Organize, OpenDirectory } from '../wailsjs/go/main/App'

let dirPath = ""

window.organize = async function () {
    if (dirPath === "") {
        alert("Selecione uma pasta")
        return
    }

    const sortBy = document.querySelector<HTMLSelectElement>('#sortBy')!.value

    await Organize(sortBy, dirPath)
        .then(() => {
            alert("Arquivos organizados com sucesso!")
        })
        .catch(err => {
            alert("Erro: " + err)
        })
}

window.openDirectory = async function () {
    dirPath = await OpenDirectory().catch(err => {
        alert("Erro: " + err)
    }) ?? ""
}

document.querySelector('#app')!.innerHTML = `
    <div class="container">
        <h1>Organizador de arquivos</h1>

        <p>1. Escolha como quer organizar seus arquivos</p>

        <select id="sortBy">
            <option value="-d">Por data de criação</option>
            <option value="-n">Por nome</option>
        </select>

        <p>2. Escolha os arquivos que deseja organizar</p>

        <button "class" onclick="openDirectory()">Abrir pasta</button>

        <button onclick="organize()">OK</button>
    </div>
`

declare global {
    interface Window {
       organize: () => void
       openDirectory: () => void
    }
}

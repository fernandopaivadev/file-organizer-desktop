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
        <p class="title">Organizador de arquivos</p>

        <p class="text">1. Escolha como quer organizar seus arquivos</p>

        <select id="sortBy">
            <option value="-d">Por dia de criação</option>
            <option value="-m">Por mês de criação</option>
            <option value="-y">Por ano de criação</option>
            <option value="-n">Por nome</option>
        </select>

        <p class="text">2. Escolha os arquivos que deseja organizar</p>

        <button class="openDirectory" onclick="openDirectory()">Abrir pasta</button>

        <button class="organize" onclick="organize()">OK</button>
    </div>
`

declare global {
    interface Window {
       organize: () => void
       openDirectory: () => void
    }
}

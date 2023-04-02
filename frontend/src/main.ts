import './style.css'
import './app.css'

import { Organize, OpenDirectory } from '../wailsjs/go/main/App'

window.toast = function (message: string) {
    const toast = document.createElement('div')
    toast.classList.add('toast')
    toast.innerText = message

    document.body.appendChild(toast)

    setTimeout(() => {
        toast.remove()
    }, 3000)
}

let dirPath = ""

window.organize = async function () {
    if (dirPath === "") {
        window.toast("Selecione uma pasta")
        return
    }

    const sortBy = document.querySelector<HTMLSelectElement>('#sortBy')!.value

    await Organize(sortBy, dirPath)
        .then(() => {
            window.toast("Arquivos organizados com sucesso!")
        })
        .catch(err => {
            window.toast("Erro: " + err)
        })
}

window.openDirectory = async function () {
    dirPath = await OpenDirectory().catch(err => {
        window.toast("Erro: " + err)
    }) ?? ""
}

document.querySelector('#app')!.innerHTML = `
    <div class="container">
        <p class="title">Organizador de arquivos</p>

        <p class="text">1. Escolha como quer organizar seus arquivos</p>

        <select id="sortBy">
            <option value="-m">Por mês de criação</option>
            <option value="-d">Por dia de criação</option>
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
       toast: (message: string) => void
    }
}

# Implementações em Go do Livro "Entendendo Algoritmos"

[![Go](https://img.shields.io/badge/Go-1.20%2B-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)

Este repositório dedica-se à implementação das estruturas de dados e algoritmos discutidos no livro "Entendendo Algoritmos" de Aditya Y. Bhargava, utilizando a linguagem de programação Go.

## Propósito

Este projeto serve como um espaço para:

* **Aprendizado de Go:** Praticar a sintaxe, os recursos e as melhores práticas da linguagem Go.
* **Compreensão de Estruturas de Dados:** Desenvolver uma compreensão prática de como as diversas estruturas de dados funcionam internamente.
* **Resolução de Exercícios:** Implementar as soluções para os exercícios propostos no livro, solidificando o conhecimento teórico.

## Conteúdo

A medida que os exercícios forem sendo resolvidos, este README será atualizado para refletir o progresso e fornecer insights sobre cada implementação. Espere encontrar:

* **Código Fonte em Go:** Implementações claras e comentadas dos algoritmos e estruturas de dados.
* **Explicações (neste README):** Uma breve descrição da minha compreensão de cada estrutura de dados implementada, incluindo seus conceitos chave e como a implementação em Go se relaciona com a teoria do livro.

## Estruturas de Dados

### Busca Binária

Algoritmo de pesquisa em listas ordenadas, que caso o elemento procurado
esteja na lista a pesquisa binária retorna sua localização, caso contrário a pesquisa binária irá retornar None.

Neste algoritmo sempre será utilizado o valor do meio para eliminar a metade de cima ou de baixo da lista para agilizar o processo de pesquisa, fazendo isso até achar ou ter certeza que o elemento desejado não está na lista.

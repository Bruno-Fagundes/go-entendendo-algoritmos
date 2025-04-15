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


## Notação Big O

A notação Big O é uma notação especial que diz o quão rápido é um algoritmo. Por exemplo, em uma lista de tamanho n o tempo de execução na notação Big O é O(n). A notação não fornece o tempo em segundos, mas sim permite que você compare o números de operações necessárias para a execução do código.

O tempo de execução na notação Big O é o O(log n). Sendo escrita, de maneira geral, da seguinte forma:

Big O ----------> O(n) <--------------------- número de operações

Isso fornece o número de operações que um algoritmo realiza. O Big O(n) é utilizada quando o algoritmo faz todas as operações para concluir uma atividade e o Big O(log n) é utilizado quando o algoritmo utiliza "atalhos", semelhante a relação da busca simples com a busca binária. O tempo de execução de um algoritmo é medido por meio de seu crescimento, sendo O(log n) mais rápido que O(n) e ficando cada vez mais conforme o número de operações aumenta. 

Suponha que você utiliza uma pesquisa simples para procurar o nome de uma pessoa em uma agenda telefônica. Você sabe que o tempo de execução O(n), o significa que na pior das hipóteses todos os nomes serão verificados. A notação Big O sempre estabelece o tempo de execução para a pior hipótese. 

Exemplos de execução Big O (ordenados do mais rápido pro mais lento):

• O(log n), também conhecido como tempo logarítmico. Exemplo: pesquisa binária.

• O(n), conhecido como tempo linear. Exemplo: pesquisa simples.

• O(n * log n). Exemplo: um algoritmo rápido de ordenação, como a ordenação quicksort.

• O(n²). Exemplo: um algoritmo lento de ordenação, como a ordenação por seleção.

• O(n!). Exemplo: um algoritmo bastante lento, como o "Problema de Caixeiro Viajante" que precisa passar cidades percorrendo uma distância mínima, mas quanto mais cidades são adicionadas maior fica o número de operações chegando a valores incalculáveis.
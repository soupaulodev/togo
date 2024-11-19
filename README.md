# TOGO - CLI Gerenciamento de Tarefas em Golang

![Gopher image](https://miro.medium.com/v2/resize:fit:735/1*_d8_TuE2kIsZnCSEamV4jA.jpeg)

## Descrição

TOGO é um gerenciador de tarefas escrito em Golang. O projeto foi criado com o intuito de aprender mais sobre a linguagem e suas funcionalidades.

## Planejamento do Projeto

### Funcionalidades do Sistema

#### 1. Criar Tarefa

- Cada tarefa terá atributos como título, descrição, status (pendente, concluída), e prioridade.
- As tarefas serão armazenadas em memória.

#### 2. Listar Tarefas

- Mostrar todas as tarefas ou filtrar por status.

#### 3. Atualizar Tarefa

- Alterar atributos de uma tarefa específica (ex.: mudar o status para "concluída").

#### 4. Excluir Tarefa

- Remover uma tarefa da lista.

#### 5. Gerenciar Concorrência

- Suporte a múltiplos usuários simultâneos com proteção de dados usando sync.Mutex.
- Suporte a múltiplos usuários simultâneos com proteção de dados usando sync.Mutex.

#### 6. Notificações Simuladas

- Um goroutine pode simular notificações para tarefas pendentes.

## Estrutura do Projeto

```css
    task-manager/
    ├── main.go               // Ponto de entrada
    ├── task/
    │   ├── task.go           // Definição da estrutura e lógica das tarefas
    │   ├── manager.go        // Gerenciamento de tarefas (CRUD)
    ├── utils/
    │   ├── validation.go     // Validação de dados
    │   └── logger.go         // Funções de logging
    └── config/
        └── config.go         // Configurações gerais
```

## Instalação

Para instalar o TOGO acesse a página de [releases](a) e baixe a versão mais recente do projeto de acordo com o seu sistema operacional.

- [Windows](a)
- [Linux](a)
- [MacOS](a)

## Desenvolvimento

### Pré-requisitos

Para desenvolver o projeto, você precisa ter o Go na versão 1.23.2 instalado em sua máquina. Para verificar se você já tem o Go instalado, execute o seguinte comando:

```sh
go version
```

Caso você não tenha o Go instalado, acesse a [página oficial](https://golang.org/dl/) e siga as instruções de instalação.

### Clonando o repositório

Para clonar o repositório, execute o seguinte comando:

```sh
git clone https://github.com/soupaulodev/togo
```

### Executando o projeto

```sh
go run main.go
```

## License

This project is licensed under the MIT License - see the [license](https://github.com/soupaulodev/togo/blob/main/LICENSE) file for details.

## Contributing

First of all, thank you for considering contributing to this project. Any help is appreciated. If you want to contribute, follow these steps:

1. Fork the project
2. Create a new branch (`git checkout -b feature/feature-name`)
3. Make the changes
4. Commit the changes
5. Push to the branch (`git push origin feature/feature-name`)
6. Create a new Pull Request

## Author

Paulo Marques - [soupaulodev](https://soupaulodev.com.br)

<!-- Passos para construir  -->

<!-- Login e Criação de Conta -->
    <!-- Validadores de campos --> ✅
    <!-- Hash de senhas --> ✅
    <!-- Geração de token JWT --> ✅
    <!-- Testes de service, handler e routes --> ✅

<!-- Construir get psiquiatras e get terapeutas --> ✅
<!-- Adicionar image para display no frontend -->
<!-- selecionar por role buscando users para retornar:
 imagem, nome, email, idade, especialidade SE terapeuta --> ✅
<!-- imagem, nome, email, idade SE psiquiatra --> ✅
<!-- adicionar coluna image na tabela users --> ✅
<!-- Pacientes podem ter apenas 1 psiquiatra e 1 terapeuta --> ✅
<!-- 1 terapeuta ou 1 psiquitra podem ter varios pacientes --> ✅

<!-- Detalhes e começo da consulta -->
    <!-- Image, Descrição, Valor da Consulta, Agenda(Horários -> Dia, Hora) -->
        <!-- Psiquiatra ou Terapeuta -> 
            Agenda Dia e Agenda Horário (reserved) tirar do display -->
                <!-- Quando clicar no botão de agendar começar consulta
            <!-- Pega o consulta_id pra indentificar e começar -->
            <!-- Final da consulta pagar o valor da consulta ao psiquiatra ou terapeuta (balance += valor da consulta) -->
            
<!-- Consulta -->
<!-- Consulta vai ter um id único -->
<!-- Usuário e Psiquiatra precisam apresentar um token para começar a consulta -->
    <!-- Se usuário for paciente -->
        <!-- Tela de Video -->
        <!-- Salvar terapeuta ou psiquiatra se gostar -->
            <!-- paciente.terapeuta -->
            <!-- paciente.psiquitra -->

    <!-- Se usuário for terapeuta -->
        <!-- Tela de vídeo -->
        <!-- Recomendar Livros -->
            <!-- Form de campo de recomendação de livro -->
        <!-- Anotações sobre paciente -->
        <!-- Encerrar consulta -->
            <!-- Salva livros(se tiver) -->
                consulta.livros += livro
            <!-- Salva anotações e  se houver -->
                consulta.anotacoes que no caso sera uma lista de anotações

    <!-- Se usuário for psiquiatra -->
        <!-- Tela de video -->
        <!-- Adicionar Remédios -->
        <!-- Opcional(Provavel diagnóstico) -->
        <!-- Anotações -->
        <!-- Encerrar consulta -->
            Salvar remédios(quantidade, nome)
            <!-- Salva anotações e diagnostico se houver -->
                consulta.diagnostico e consulta.anotacoes que no caso sera uma lista de anotações
            
<!-- Dados do paciente -->
<!-- Get todos pacientes do terapeuta ou psiquitra -->
    <!-- Nome, Imagem, Email -->
    <!-- Botão detalhes -->
    <!-- Se terapeuta -->
        <!-- Pegar paciente com um get id -->
            <!-- mostrar  anotações do paciente -->
            <!-- Livros recomendado -->
    <!-- Se psiquiatra -->
        <!-- Pegar paciente com um get id -->
            <!-- mostrar diagnosticos -->
            <!-- Remédios(quantidade, nome) -->
            <!-- mostrar anotações do paciente -->

<!-- Preciso criar um update de remédios -->
<!-- E ou Create -->
<!-- Vai ser uma lista de remédios(nome, quantidade) -->
<!-- Psquiatra pode alterar essas informações -->
<!-- Ou deletar o remédio -->

<!-- Perfil do paciente-->
    <!-- Get do seu psiquiatra pelo id -->
    <!-- Get do seu terapeuta pelo id -->
        <!-- Anotações do psiquiatra e do terapeuta -->
        <!-- Seus remédios -->
        <!-- Seu diagnóstico -->

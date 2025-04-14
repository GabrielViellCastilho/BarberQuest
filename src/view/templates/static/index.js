document.addEventListener("DOMContentLoaded", function () {
    fetch("/static/footer.html")
        .then(response => response.text())
        .then(html => {
            document.getElementById("footer-container").innerHTML = html;
        })
        .catch(error => console.error("Erro ao carregar o footer:", error));
});

document.addEventListener("DOMContentLoaded", function () {
    fetch("/static/modals.html")
        .then(response => response.text())
        .then(html => {
            document.getElementById("modals-placeholder").innerHTML = html;

            // Inicializa os event listeners após o carregamento dos modals
            initializeLoginForm();
            initializeCreateAccountForm();
            initializeForgotPasswordForm();
        })
        .catch(error => console.error("Erro ao carregar modals:", error));
});

// Função para inicializar o formulário de login
function initializeLoginForm() {
    const loginForm = document.getElementById("Forms_Login");

    if (loginForm) {
        loginForm.addEventListener("submit", async function(event) {
            event.preventDefault(); // Impede o envio automático do formulário

            const email = document.getElementById("email").value;
            const password = document.getElementById("password").value;

            const userData = {
                email: email,
                password: password
            };

            try {
                const response = await fetch("/login", {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json"
                    },
                    body: JSON.stringify(userData)
                });

                console.log("Headers recebidos:", response.headers);

                if (!response.ok) {
                    if (response.status == 404 || response.status == 400){
                        throw new Error("Email ou senha incorreta")
                    }
                    let errorMessage = `Erro ${response.status}: ${response.statusText}`;
                    console.log(errorMessage)
                    throw new Error(errorMessage);
                }

                const authHeader = response.headers.get("Authorization");

                if (authHeader) {
                    const token = authHeader.startsWith("Bearer ") ? authHeader.split(" ")[1] : authHeader;
                    localStorage.setItem("jwt_token", token);

                    showAlert("Login realizado com sucesso")
                    console.log(localStorage.getItem("jwt_token"))
                } else {
                    throw new Error("Token JWT não encontrado no header.");
                }

                let modalElement = document.getElementById("modalLogin");
                let modal = bootstrap.Modal.getInstance(modalElement);
                modal.hide();

                setTimeout(() => {
                    window.location.reload();
                }, 3000);
            } catch (error) {
                console.error("Erro no login:", error);
                showAlert(error.message);
            }
        });
    } else {
        console.error("Elemento #Forms_Login não encontrado!");
    }
}

// Função para inicializar o formulário de criação de conta
function initializeCreateAccountForm() {
    const createForm = document.getElementById("Forms_Novo_Usuario");

    if (createForm) {
        createForm.addEventListener("submit", async function(event) {
            event.preventDefault(); // Impede o envio automático do formulário

            const name = document.getElementById("createUserName").value;
            const email = document.getElementById("createUserEmail").value;
            const cellphone = document.getElementById("createUserPhone").value;
            const dateOfBirth = document.getElementById("createUserDateOfBirth").value;
            const password = document.getElementById("createUserPassword").value;
            const confirmPassword = document.getElementById("createUserConfirmPassword").value;

            // 🔹 Validação da data de nascimento
            if (!dateOfBirth) {
                return showAlert("Por favor, preencha a data de nascimento.");
            }

            const today = new Date();
            const selectedDate = new Date(dateOfBirth);

            // 🔹 Verifica se a data de nascimento não está no futuro
            if (selectedDate > today) {
                return showAlert("A data de nascimento não pode ser no futuro.");
            }

            // 🔹 Formatar a data para YYYY-MM-DD
            const formattedDateOfBirth = selectedDate.toISOString().split("T")[0];

            // 🔹 Validação de senha
            const passwordRegex = /^(?=.*[A-Za-z])(?=.*\d)(?=.*[!@#$%*]).{6,}$/;
            if (!passwordRegex.test(password)) {
                return showAlert("A senha deve ter pelo menos 6 caracteres, conter letras, números e um caractere especial (!@#$%*).");
            }

            if (password !== confirmPassword) {
                return showAlert("As senhas não coincidem.");
            }

            const userCreateData = {
                name: name,
                email: email,
                password: password,
                cellphone: cellphone.replace(/\D/g, ''), // Remover a máscara
                dateOfBirth: formattedDateOfBirth // Adicionando a data formatada
            };

            try {
                console.log("Enviando a requisição");
                const response = await fetch("/createCustomerUser", {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json"
                    },
                    body: JSON.stringify(userCreateData)
                });
                console.log("Recebendo");

                if (!response.ok) {
                    if (response.status.toString() === "409") {
                        showAlert("Email já está cadastrado");
                        return;
                    }
                    let errorMessage = `Erro ${response.status}: ${response.statusText}`;
                    throw new Error(errorMessage);
                }

                showAlert("Conta criada com sucesso!");

                // 🔹 Obtendo o token do header Authorization
                const authHeader = response.headers.get("Authorization");

                if (authHeader) {
                    const token = authHeader.startsWith("Bearer ") ? authHeader.split(" ")[1] : authHeader;

                    // 🔹 Salvando o token no LocalStorage
                    localStorage.setItem("jwt_token", token);

                } else {
                    throw new Error("Token JWT não encontrado no header.");
                }

                // Redirecionamento após 3 segundos
                setTimeout(() => {
                    window.location.reload(); // Recarrega a página
                }, 3000);

            } catch (error) {
                console.error("Erro ao criar conta:", error);
                showAlert("Erro ao criar conta. Tente novamente mais tarde.");
            }
        });
    } else {
        console.error("Elemento #Forms_Novo_Usuario não encontrado!");
    }
}

// Função para inicializar o formulário de "Esqueci a senha"
function initializeForgotPasswordForm() {
    const changePasswordForm = document.getElementById("FormsChangePassword");

    if (changePasswordForm) {
        changePasswordForm.addEventListener("submit", async function(event) {
            event.preventDefault(); // Impede o envio automático do formulário

            const email = document.getElementById("emailChangePassword").value;

            try {
                // 🔹 Envia o e-mail como parte da URL
                const response = await fetch(`/sendEmailForgotPassword/${encodeURIComponent(email)}`, {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json"
                    }
                });

                if (!response.ok) {
                    let errorMessage = `Erro ${response.status}: ${response.statusText}`;
                    console.error(errorMessage);
                    throw new Error(errorMessage);
                }

                // 🔹 Fecha o modal corretamente
                let modalElement = document.getElementById("modalChangePassword");
                let modal = bootstrap.Modal.getInstance(modalElement);
                modal.hide();

                showAlert("Verifique seu email.");

            } catch (error) {
                console.error("Erro ao enviar email:", error);
                showAlert(error.message);
            }
        });
    } else {
        console.error("Elemento #FormsChangePassword não encontrado!");
    }
}


document.addEventListener("DOMContentLoaded", updateUI);

function updateUI() {
    const token = localStorage.getItem("jwt_token");
    const loginButton = document.getElementById("loginButton");
    const profileButton = document.getElementById("profileButton");

    if (token) {
        // Se o usuário estiver logado, exibe o botão de perfil e oculta o de login
        loginButton.style.display = "none";
        profileButton.style.display = "block";
    } else {
        // Se não estiver logado, exibe o botão de login e oculta o de perfil
        loginButton.style.display = "block";
        profileButton.style.display = "none";
    }
}

function logout() {
    localStorage.removeItem("jwt_token"); // Remove o token
    window.location.href = "/";
}

document.addEventListener("DOMContentLoaded", function () {
    flatpickr("#datePicker", {
        dateFormat: "d/m/Y", // Formato: Dia/Mês/Ano
        minDate: "today", // Impede datas passadas
        locale: "pt", // Define o calendário para português
        disable: [
            function(date) {
                return date.getDay() === 0; // Desabilita os domingos (0 = Domingo)
            }
        ],
        disableMobile: true // Força exibição no mobile
    });
});

// Função para formatar o número de telefone
function formatPhoneNumber() {
    var phoneInput = document.getElementById('createUserPhone');
    var myForm = document.forms.myForm;
    var result = document.getElementById('result');  // Apenas para fins de depuração

    // Máscara de telefone
    phoneInput.addEventListener('input', function (e) {
        // Remove caracteres não numéricos e aplica a máscara de telefone brasileira
        var x = e.target.value.replace(/\D/g, '').match(/(\d{0,2})(\d{0,5})(\d{0,4})/);
        e.target.value = !x[2] ? x[1] : '(' + x[1] + ') ' + x[2] + (x[3] ? '-' + x[3] : '');
    });

    // Quando o formulário for enviado, limpar a máscara e mostrar o número apenas
    myForm.addEventListener('submit', function(e) {
        // Remove todos os caracteres não numéricos antes de enviar
        phoneInput.value = phoneInput.value.replace(/\D/g, '');
        result.innerText = phoneInput.value;  // Apenas para fins de depuração

        e.preventDefault(); // Evitar o envio do formulário para teste
    });
}

// Chama a função para formatar o telefone quando a página carrega
document.addEventListener('DOMContentLoaded', function () {
    formatPhoneNumber();
});

// Função para exibir alerta
function showAlert(message, type = "info") {
    const alertTitle = document.getElementById("alertTitle");
    const alertMessage = document.getElementById("alertMessage");

    // Define o título e ícone do modal conforme o tipo de alerta
    if (type === "success") {
        alertTitle.innerHTML = "✅ Sucesso";
    } else if (type === "error") {
        alertTitle.innerHTML = "❌ Erro";
    } else {
        alertTitle.innerHTML = "";
    }

    // Define a mensagem no modal e mantém a formatação
    alertMessage.innerHTML = message.replace(/\n/g, "<br>");

    // Fecha o modal ativo antes de abrir o alerta
    const activeModal = document.querySelector(".modal.show");
    if (activeModal) {
        let modalInstance = bootstrap.Modal.getInstance(activeModal);
        modalInstance.hide();
    }

    // Pequeno atraso para evitar sobreposição de modais
    setTimeout(() => {
        let modalAlert = new bootstrap.Modal(document.getElementById("modalAlert"));
        modalAlert.show();
    }, 300);
}

function showTermsOfUse() {
    showAlert(`
        <strong>Termos de Uso</strong><br><br>

        <strong>1. Introdução</strong><br>
        Bem-vindo ao site de agendamento da Spartan Barbearia. Ao criar uma conta e utilizar nossos serviços, você concorda com os termos e condições estabelecidos abaixo.<br><br>

        <strong>2. Cadastro e Conta do Usuário</strong><br>
        Para utilizar nossos serviços, é necessário criar uma conta fornecendo informações precisas e atualizadas.<br>
        As contas podem ser de três tipos: usuário, barbeiro e administrador.<br>
        Você é responsável por manter a confidencialidade de suas credenciais de acesso.<br>
        Reservamo-nos o direito de suspender ou excluir contas que violem estes termos.<br><br>

        <strong>3. Agendamentos</strong><br>
        Os agendamentos são realizados por meio da plataforma e estão sujeitos à disponibilidade.<br>
        Cancelamentos e reagendamentos podem ser feitos a qualquer momento.<br>
        O pagamento dos serviços deve ser realizado presencialmente na barbearia.<br><br>

        <strong>4. Responsabilidades do Usuário</strong><br>
        Utilizar a plataforma de forma lícita e respeitosa.<br>
        Não realizar agendamentos falsos ou fraudulentos.<br>
        Respeitar os profissionais e as políticas da barbearia.<br><br>

        <strong>5. Modificações nos Termos de Uso</strong><br>
        Podemos atualizar estes termos periodicamente, e as alterações entrarão em vigor a partir da data de publicação.<br>
        É sua responsabilidade revisar os termos regularmente.<br><br>

        <strong>6. Privacidade e Proteção de Dados</strong><br>
        Os dados coletados no cadastro incluem nome completo, e-mail, telefone, data de nascimento e senha.<br>
        A data de nascimento pode ser utilizada para fins de promoções.<br>
        O telefone pode ser utilizado para contato relacionado aos agendamentos.<br>
        Não compartilhamos suas informações com terceiros sem sua autorização.<br><br>

        <strong>7. Contato</strong><br>
        Para esclarecimentos ou dúvidas sobre estes termos, entre em contato pelo telefone +55 (12) 98803-313.<br>
    `);
}


function showConfirm(title, message, onConfirm) {
    // Configura o título e a mensagem do modal
    document.getElementById("confirmTitle").textContent = title;
    document.getElementById("confirmMessage").textContent = message;

    // Mostrar o modal
    const modal = new bootstrap.Modal(document.getElementById('modalConfirm'));
    modal.show();

    // Ação para o botão "Confirmar"
    document.getElementById("confirmButton").addEventListener("click", function() {
        onConfirm();  // Executa a ação de confirmação passada como parâmetro
        modal.hide(); // Fecha o modal
    });

    // Ação para o botão "Cancelar"
    document.getElementById("cancelButton").addEventListener("click", function() {
        modal.hide(); // Apenas fecha o modal
    });
}

function isTokenExpired() {
    const token = localStorage.getItem("jwt_token");
    if (!token){
        return true
    }

    try {
        const payload = JSON.parse(atob(token.split(".")[1])); // Decodifica o payload do token
        const exp = payload.exp * 1000; // Converte a expiração para milissegundos

        // Verifica se a data de expiração do token já passou
        if (Date.now() >= exp) {
            showAlert("Sessão expirada, faça login novamente");
            setTimeout(() => {
                logout();  // Chama a função de logout após 3 segundos
            }, 3000);
            return true; // Se expirou, retorna true
        }

        return false; // Caso o token ainda seja válido, retorna false
    } catch (error) {
        console.error("Erro ao decodificar token:", error);
        showAlert("Erro ao verificar o token");
        return true; // Se houver erro ao decodificar o token, considera que a sessão expirou
    }
}


function getUserRole() {
    const token = localStorage.getItem("jwt_token");
    if (!token) return null;

    try {
        const base64Url = token.split('.')[1]; // Pega a parte do payload do JWT
        const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
        const jsonPayload = decodeURIComponent(atob(base64).split('').map(c => '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2)).join(''));

        const decoded = JSON.parse(jsonPayload);
        return decoded.role; // Retorna a role
    } catch (error) {
        console.error("Erro ao decodificar token:", error);
        return null;
    }
}

document.addEventListener("DOMContentLoaded", function () {
    const role = getUserRole();
    const profileButton = document.getElementById("profileButton");

    if (!profileButton) return;

    let offcanvasId = "";

    if (role === "admin") {
        offcanvasId = "offcanvasAdmin";
    } else if (role === "barber") {
        offcanvasId = "offcanvasBarber";
    } else if (role=="user"){
        offcanvasId = "offcanvasUser";
    }

    profileButton.setAttribute("data-bs-target", `#${offcanvasId}`);
});

function showConfirm(title, message, onConfirm) {
    document.getElementById("confirmTitle").textContent = title;
    document.getElementById("confirmMessage").textContent = message;

    const modal = new bootstrap.Modal(document.getElementById('modalConfirm'));
    modal.show();

    document.getElementById("confirmButton").onclick = function () {
        onConfirm();
        modal.hide();
    };

    document.getElementById("cancelButton").onclick = function () {
        modal.hide();
    };
}

function getOnlyNumbers(phone) {
    return phone.replace(/\D/g, '');
}

function formatCellphoneInput(cellphone) {
    // Remove todos os caracteres que não são números
    cellphone = cellphone.replace(/\D/g, '');

    // Aplica a formatação do celular brasileiro
    if (cellphone.length > 2) {
        cellphone = `(${cellphone.substring(0, 2)}) ${cellphone.substring(2)}`;
    }
    if (cellphone.length > 10) {
        cellphone = `${cellphone.substring(0, 10)}-${cellphone.substring(10, 15)}`;
    }

    return cellphone;
}


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

            initializeLoginForm();
            initializeCreateAccountForm();
            initializeForgotPasswordForm();
        })
        .catch(error => console.error("Erro ao carregar modals:", error));
});

function initializeLoginForm() {
    const loginForm = document.getElementById("Forms_Login");

    if (loginForm) {
        loginForm.addEventListener("submit", async function(event) {
            event.preventDefault();

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

function initializeCreateAccountForm() {
    const createForm = document.getElementById("Forms_Novo_Usuario");

    if (createForm) {
        createForm.addEventListener("submit", async function(event) {
            event.preventDefault();

            const name = document.getElementById("createUserName").value;
            const email = document.getElementById("createUserEmail").value;
            const cellphone = document.getElementById("createUserPhone").value;
            const dateOfBirth = document.getElementById("createUserDateOfBirth").value;
            const password = document.getElementById("createUserPassword").value;
            const confirmPassword = document.getElementById("createUserConfirmPassword").value;

            if (!dateOfBirth) {
                return showAlert("Por favor, preencha a data de nascimento.");
            }

            const today = new Date();
            const selectedDate = new Date(dateOfBirth);

            if (selectedDate > today) {
                return showAlert("A data de nascimento não pode ser no futuro.");
            }

            const formattedDateOfBirth = selectedDate.toISOString().split("T")[0];

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
                cellphone: cellphone.replace(/\D/g, ''),
                dateOfBirth: formattedDateOfBirth
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


                const authHeader = response.headers.get("Authorization");

                if (authHeader) {
                    const token = authHeader.startsWith("Bearer ") ? authHeader.split(" ")[1] : authHeader;


                    localStorage.setItem("jwt_token", token);

                } else {
                    throw new Error("Token JWT não encontrado no header.");
                }


                setTimeout(() => {
                    window.location.reload();
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


function initializeForgotPasswordForm() {
    const changePasswordForm = document.getElementById("FormsChangePassword");

    if (changePasswordForm) {
        changePasswordForm.addEventListener("submit", async function(event) {
            event.preventDefault();

            const email = document.getElementById("emailChangePassword").value;

            try {

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

        loginButton.style.display = "none";
        profileButton.style.display = "block";
    } else {
        loginButton.style.display = "block";
        profileButton.style.display = "none";
    }
}

function logout() {
    localStorage.removeItem("jwt_token");
    window.location.href = "/";
}

document.addEventListener("DOMContentLoaded", function () {
    flatpickr("#datePicker", {
        dateFormat: "d/m/Y",
        minDate: "today",
        locale: "pt",
        disable: [
            function(date) {
                return date.getDay() === 0;
            }
        ],
        disableMobile: true
    });
});

function formatPhoneNumber() {
    var phoneInput = document.getElementById('createUserPhone');
    var myForm = document.forms.myForm;
    var result = document.getElementById('result');


    phoneInput.addEventListener('input', function (e) {

        var x = e.target.value.replace(/\D/g, '').match(/(\d{0,2})(\d{0,5})(\d{0,4})/);
        e.target.value = !x[2] ? x[1] : '(' + x[1] + ') ' + x[2] + (x[3] ? '-' + x[3] : '');
    });


    myForm.addEventListener('submit', function(e) {

        phoneInput.value = phoneInput.value.replace(/\D/g, '');
        result.innerText = phoneInput.value;

        e.preventDefault();
    });
}


document.addEventListener('DOMContentLoaded', function () {
    formatPhoneNumber();
});


function showAlert(message, type = "info") {
    const alertTitle = document.getElementById("alertTitle");
    const alertMessage = document.getElementById("alertMessage");


    if (type === "success") {
        alertTitle.innerHTML = "✅ Sucesso";
    } else if (type === "error") {
        alertTitle.innerHTML = "❌ Erro";
    } else {
        alertTitle.innerHTML = "";
    }

    alertMessage.innerHTML = message.replace(/\n/g, "<br>");


    const activeModal = document.querySelector(".modal.show");
    if (activeModal) {
        let modalInstance = bootstrap.Modal.getInstance(activeModal);
        modalInstance.hide();
    }


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

    document.getElementById("confirmTitle").textContent = title;
    document.getElementById("confirmMessage").textContent = message;


    const modal = new bootstrap.Modal(document.getElementById('modalConfirm'));
    modal.show();


    document.getElementById("confirmButton").addEventListener("click", function() {
        onConfirm();
        modal.hide();
    });


    document.getElementById("cancelButton").addEventListener("click", function() {
        modal.hide();
    });
}

function isTokenExpired() {
    const token = localStorage.getItem("jwt_token");
    if (!token){
        return true
    }

    try {
        const payload = JSON.parse(atob(token.split(".")[1]));
        const exp = payload.exp * 1000;

        if (Date.now() >= exp) {
            showAlert("Sessão expirada, faça login novamente");
            setTimeout(() => {
                logout();
            }, 3000);
            return true;
        }

        return false;
    } catch (error) {
        console.error("Erro ao decodificar token:", error);
        showAlert("Erro ao verificar o token");
        return true;
    }
}


function getUserRole() {
    const token = localStorage.getItem("jwt_token");
    if (!token) return null;

    try {
        const base64Url = token.split('.')[1];
        const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
        const jsonPayload = decodeURIComponent(atob(base64).split('').map(c => '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2)).join(''));

        const decoded = JSON.parse(jsonPayload);
        return decoded.role;
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
    cellphone = cellphone.replace(/\D/g, '');


    if (cellphone.length > 2) {
        cellphone = `(${cellphone.substring(0, 2)}) ${cellphone.substring(2)}`;
    }
    if (cellphone.length > 10) {
        cellphone = `${cellphone.substring(0, 10)}-${cellphone.substring(10, 15)}`;
    }

    return cellphone;
}


const form = document.getElementById('login-form');

form.addEventListener('submit', (event) => {
  event.preventDefault();

  const codigo = document.getElementById('codigo').value;
  console.log("Código ingresado:", codigo); // Mostrar el código ingresado


  fetch('/api/v1/Usuarios/validate', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({ codigo: parseInt(codigo) }) // Convertir a número
  })
  .then(response => {
    console.log("Respuesta de la API:", response); // Mostrar la respuesta de la API
    if (response.ok) {
      window.location.href = '/menu';
    } else {
      alert('Código inválido');
    }
  })
  .catch(error => {
    console.error('Error al validar el código:', error);
    alert('Error al validar el código');
  });
});
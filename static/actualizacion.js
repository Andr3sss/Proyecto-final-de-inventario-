document.addEventListener('DOMContentLoaded', (event) => {
    fetch('/api/v1/productos/')
      .then(response => response.json())
      .then(productos => {
        const selectProducto = document.getElementById('producto');
        productos.forEach(producto => {
          const option = document.createElement('option');
          option.value = producto.ID;
          option.textContent = producto.Nombre;
          selectProducto.appendChild(option);
        });
      })
      .catch(error => console.error('Error al obtener productos:', error));

  // Obtener la lista de categorías
  fetch('/api/v1/categorias')
    .then(response => {
      if (!response.ok) {
        throw new Error('Error al obtener categorías');
      }
      return response.json();
    })
    .then(categorias => {
        const selectCategoria = document.getElementById('categoria');
        categorias.forEach(categoria => {
          const option = document.createElement('option');
          option.value = categoria.ID;
          option.textContent = categoria.Nombre;
          selectCategoria.appendChild(option);
        });
    })
    .catch(error => console.error('Error al obtener categorías:', error));
  
// Manejar el envío del formulario de productos
const formProducto = document.getElementById('formulario-producto');

formProducto.addEventListener('submit', (event) => {
  event.preventDefault();
  
  const productoID = document.getElementById('producto').value; // Obtener el ID del producto
  const descripcion = document.getElementById('descripcion').value;
  const precio = parseFloat(document.getElementById('precio').value);
  const color = document.getElementById('color').value;
  const categoriaID = parseInt(document.getElementById('categoria').value);



  const productoActualizado = {
    prducto_id: productoID,
    descripcion: descripcion,
    precio: precio,
    color: color,
    categoria_id: categoriaID
  };

  console.log("ID del producto:", productoID); // Mostrar el ID del producto
  console.log("Datos del producto:", productoActualizado); // Mostrar los datos del producto

  fetch(`/api/v1/productos/${productoID}`, { // Ruta corregida
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(productoActualizado)
  })
  .then(response => {
    console.log("Respuesta de la API:", response); // Mostrar la respuesta de la API


    if (response.ok) {
      alert('Actualizacion exitosa');
      formProducto.reset();
      location.reload();
    } else {
      alert('Error al actualizar el producto');
    }
  })
  .catch(error => console.error('Error al agregar producto:', error));
});
  });
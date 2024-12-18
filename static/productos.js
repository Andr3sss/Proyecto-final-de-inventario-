document.addEventListener('DOMContentLoaded', (event) => {
// Obtener la lista de productos
fetch('/api/v1/productos/')
  .then(response => response.json())
  .then(productos => {
    const listaProductos = document.getElementById('lista-productos');
    productos.forEach(producto => {
      const li = document.createElement('li');
      li.textContent = `${producto.nombre} - ${producto.Categoria.nombre} - $${producto.precio} - Stock: ${producto.stock}`;
      listaProductos.appendChild(li);
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

    const nombre = document.getElementById('nombre').value;
    const descripcion = document.getElementById('descripcion').value;
    const precio = parseFloat(document.getElementById('precio').value);
    const color = document.getElementById('color').value;
    const categoriaID = parseInt(document.getElementById('categoria').value);

    const nuevoProducto = {
      nombre: nombre,
      descripcion: descripcion,
      precio: precio,
      color: color,
      categoria_id: categoriaID
    };

    fetch('/api/v1/productos/', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(nuevoProducto)
    })
    .then(response => {
      if (response.ok) {
        alert('Producto agregado correctamente');
        formProducto.reset();
        location.reload();
      } else {
        alert('Error al agregar producto');
      }
    })
    .catch(error => console.error('Error al agregar producto:', error));
  });
});
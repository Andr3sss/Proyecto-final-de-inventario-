document.addEventListener('DOMContentLoaded', (event) => {
    // Obtener la lista de productos para el select
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
  
    // Obtener la lista de movimientos
    fetch(`/api/v1/movimientos/`)
      .then(response => response.json())
      .then(movimientos => {
        const tablaMovimientos = document.getElementById('tabla-movimientos').getElementsByTagName('tbody')[0];
  
        movimientos.forEach(movimiento => {
          const row = tablaMovimientos.insertRow();
          const idCell = row.insertCell();
          const productoCell = row.insertCell();
          const tipoCell = row.insertCell();
          const cantidadCell = row.insertCell();
          const fechaCell = row.insertCell();
          const motivoCell = row.insertCell();
          const accionesCell = row.insertCell();
  
          idCell.textContent = movimiento.id;
          productoCell.textContent = movimiento.ProductoID; // Aquí deberías mostrar el nombre del producto
          tipoCell.textContent = movimiento.tipo_movimiento;
          cantidadCell.textContent = movimiento.cantidad;
          fechaCell.textContent = movimiento.fecha;
          motivoCell.textContent = movimiento.motivo;
  
          // Agregar botones de acciones (editar, eliminar)
          accionesCell.innerHTML = `
            <button class="btn-editar" data-movimiento-id="${movimiento.id}">Editar</button>
            <button class="btn-eliminar" data-movimiento-id="${movimiento.id}">Eliminar</button>
          `;
        });
  
        // Manejar clic en botones de eliminar
        const botonesEliminar = document.querySelectorAll('.btn-eliminar');
        botonesEliminar.forEach(boton => {
          boton.addEventListener('click', (event) => {
            const movimientoID = event.target.dataset.movimientoId;
            if (confirm(`¿Estás seguro de que quieres eliminar el movimiento ${movimientoID}?`)) {
              fetch(`/api/v1/movimientos/${movimientoID}`, {
                method: 'DELETE'
              })
              .then(response => {
                if (response.ok) {
                  alert('Movimiento eliminado correctamente');
                  location.reload(); // Recargar la página para actualizar la lista
                } else {
                  alert('Error al eliminar movimiento');
                }
              })
              .catch(error => console.error('Error al eliminar movimiento:', error));
            }
          });
        });
      })
      .catch(error => console.error('Error al obtener movimientos:', error));
  
    // Manejar el envío del formulario de movimientos
    const formMovimiento = document.getElementById('formulario-movimiento');
    formMovimiento.addEventListener('submit', (event) => {
      event.preventDefault();

      const selectProducto = document.getElementById('producto');
      const productoID = selectProducto.value;
      const tipoMovimiento = document.querySelector('input[name="tipo_movimiento"]:checked').value;
      const cantidad = parseInt(document.getElementById('cantidad').value);
      const motivo = document.getElementById('motivo').value;
  
      const nuevoMovimiento = {
        producto_id: productoID,
        tipo_movimiento: tipoMovimiento,
        cantidad: cantidad,
        motivo: motivo
      };
  
      fetch(`/api/v1/movimientos/${productoID}`, { // Ajustar la ruta si es necesario
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(nuevoMovimiento)
      })
      .then(response => {
        if (response.ok) {
          alert('Movimiento registrado correctamente');
          formMovimiento.reset();

              // Actualizar la tabla de inventario
    const productoID = parseInt(document.getElementById('producto').value);
    const tablaInventario = document.getElementById('tabla-inventario').getElementsByTagName('tbody')[0];
    const filas = tablaInventario.getElementsByTagName('tr');
    for (let i = 0; i < filas.length; i++) {
      const fila = filas[i];
      const celdas = fila.getElementsByTagName('td');
      if (celdas[0].textContent == productoID) { // Buscar la fila del producto
        const stockCell = celdas[3];
        const stockActual = parseInt(stockCell.textContent);
        const nuevoStock = tipoMovimiento === 'entrada' ? stockActual + cantidad : stockActual - cantidad;
        stockCell.textContent = nuevoStock;
        break;
      }
    } 
        } else {
          alert('Error al registrar movimiento');
        }
      })
      .catch(error => console.error('Error al registrar movimiento:', error));
    });
  });
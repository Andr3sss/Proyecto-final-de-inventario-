document.addEventListener('DOMContentLoaded', () => {
    const tablaInventario = document.getElementById('tabla-inventario').getElementsByTagName('tbody')[0];

    const actualizarStock = (movimientos) => {
        let stockActual = 0;
    
        // Si no hay movimientos, el stock es 0
        if (!movimientos || movimientos.length === 0) {
            return stockActual;
        }
    
        // Calcular stock basado en movimientos
        movimientos.forEach(movimiento => {
            if (movimiento.tipo_movimiento === 'entrada') {
                stockActual += movimiento.cantidad;
            } else if (movimiento.tipo_movimiento === 'salida') {
                stockActual -= movimiento.cantidad;
            }
        });
        return stockActual;
    };
    
    const renderTabla = (productos) => {
        tablaInventario.innerHTML = ''; 
        productos.forEach(producto => {
            try {
                // Calcular el stock actual
                const stockActual = producto.Stock;  
                productos.forEach(producto => {
                    const stockActual = actualizarStock(producto.Movimientos);
                    console.log("Stock calculado para el producto:", producto.Nombre, "Stock:", stockActual); 
                });

                const row = tablaInventario.insertRow();
                const nombreCell = row.insertCell();
                const categoriaCell = row.insertCell();
                const colorCell = row.insertCell();
                const stockCell = row.insertCell(); 
                const precioCell = row.insertCell();

                stockCell.textContent = stockActual; 
                nombreCell.textContent = producto.Nombre;
                categoriaCell.textContent = producto.Categoria ? producto.Categoria.Nombre : 'Sin categoría';
                colorCell.textContent = producto.Color || 'Sin color';
                stockCell.textContent = stockActual;
                precioCell.textContent = producto.Precio || 0;
            } catch (error) {
                console.error("Error al procesar el producto", producto, ":", error);
            }
        });
    };


    fetch('/api/v1/productos/')
    .then(response => response.json())
    .then(data => {
      data.forEach(producto => {
        // Aquí se asume que los productos incluyen los movimientos ahora
        let row = document.createElement('tr');
        row.innerHTML = `
          <td>${producto.Nombre}</td>
          <td>${producto.Categoria.Nombre}</td>
          <td>${producto.Color}</td>
          <td>${producto.Stock}</td> <!-- Mostrar el stock calculado -->
          <td>${producto.Precio}</td>
        `;
        document.querySelector('#tabla-inventario tbody').appendChild(row);
      });
    })
    .catch(error => console.error('Error al cargar los productos:', error));

});

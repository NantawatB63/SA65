import React, { useState, useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import { StocksInterface } from "../interfaces/IStock";
import { GetStock } from "../services/HttpClientService";

function Products() {
  const [stocks, setStocks] = useState<StocksInterface[]>([]);

  useEffect(() => {
    getStocks();
  }, []);

  const getStocks = async () => {
    let res = await GetStock();
    if (res) {
      setStocks(res);
    } 
  };

  const columns: GridColDef[] = [
    { field: "ID", headerName: "ID", width: 70 },
    {
      field: "Product_name",
      headerName: "Name",
      width: 200,
      valueFormatter: (params) => params.value.Stock_name,
    },
    {
        field: "Quantity",
        headerName: "Quantity",
        width: 100,
        valueFormatter: (params) => params.value.Quantity,
    },
    {
      field: "Shelf",
      headerName: "Shelf",
      width: 150,
      valueFormatter: (params) => params.value.Shelf_name,
    },
    {
      field: "Lot",
      headerName: "Lot",
      width: 150,
      valueFormatter: (params) => params.value.Lot_number,
    },
    { field: "Employee", 
      headerName: "Employee", 
      width: 150 ,
      valueFormatter: (params) => params.value.FirstName + " " + params.value.LastName,
    },

  ];

  return (
    <div>
      <Container maxWidth="md">
        <Box
          display="flex"
          sx={{
            marginTop: 2,
          }}
        >
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
             Stock Information
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/stock/create"
              variant="contained"
              color="primary"
            >
              Create Data
            </Button>
          </Box>
        </Box>
        <div style={{ height: 400, width: "100%", marginTop: "20px" }}>
          <DataGrid
            rows={stocks}
            getRowId={(row) => row.ID}
            columns={columns}
            pageSize={5}
            rowsPerPageOptions={[5]}
          />
        </div>
      </Container>
    </div>
  );
}

export default Products;
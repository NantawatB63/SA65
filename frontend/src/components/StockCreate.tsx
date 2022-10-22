import React, { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import Button from "@mui/material/Button";
import FormControl from "@mui/material/FormControl";
import Container from "@mui/material/Container";
import Paper from "@mui/material/Paper";
import Grid from "@mui/material/Grid";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import Divider from "@mui/material/Divider";
import Snackbar from "@mui/material/Snackbar";
import Select, { SelectChangeEvent } from "@mui/material/Select";
import MuiAlert, { AlertProps } from "@mui/material/Alert";
import TextField from "@mui/material/TextField";
import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { DatePicker } from "@mui/x-date-pickers/DatePicker";

import { LotsInterface } from "../interfaces/ILot";
import { ShelfsInterface } from "../interfaces/IShelf";
import { ProductsInterface } from "../interfaces/IProduct";
import { StocksInterface } from "../interfaces/IStock";

import {
  GetProducts,
  GetLot,
  GetShelf,
  Stocks,
} from "../services/HttpClientService";
const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
  props,
  ref
) {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function StockCreate() {
  const [lots, setLots] = useState<LotsInterface>();
  const [shelfs, setShelfs] = useState<ShelfsInterface>();
  const [products, setProducts] = useState<ProductsInterface[]>([]);
  const [stock, setStocks] = useState<StocksInterface>({
});

  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);

  const handleInputChange = (
    event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof StockCreate;
    const { value } = event.target;
    setStocks({ ...stock, [id]: value });
  };

  const handleClose = (
    event?: React.SyntheticEvent | Event,
    reason?: string
  ) => {
    if (reason === "clickaway") {
      return;
    }
    setSuccess(false);
    setError(false);
  };

  const handleChange = (event: SelectChangeEvent) => {
    const name = event.target.name as keyof typeof stock;
    setStocks({
      ...stock,
      [name]: event.target.value,
    });
  };

  const getProducts = async () => {
    let res = await GetProducts();
    if (res) {
      setProducts(res);
    }
  };

  const getLot = async () => {
    let res = await GetLot();
    if (res) {
        setLots(res);
    }
  };

  const getShelf = async () => {
    let res = await GetShelf();
    if (res) {
      setShelfs(res);
    }
  };

  useEffect(() => {
    getShelf();
    getLot();
    getProducts();
  }, []);

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  async function submit() {
    let data = {
      ProductID: convertType(stock.ProductID),
      Quantity: stock.Quanitiy?? "", 
      LotID: convertType(stock.LotID),
      ShelfID: convertType(stock.ShelfID),
    };

    let res = await Stocks(data);
    if (res) {
      setSuccess(true);
    } else {
      setError(true);
    }
  }

  return (
    <Container maxWidth="md">
      <Snackbar
        open={success}
        autoHideDuration={3000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "top", horizontal: "center" }}
      >
        <Alert onClose={handleClose} severity="success">
          บันทึกข้อมูลสำเร็จ
        </Alert>
      </Snackbar>
      <Snackbar
        open={error}
        autoHideDuration={6000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "top", horizontal: "center" }}
      >
        <Alert onClose={handleClose} severity="error">
          บันทึกข้อมูลไม่สำเร็จ
        </Alert>
      </Snackbar>
      <Paper>
        <Box
          display="flex"
          sx={{
            marginTop: 2,
          }}
        >
          <Box sx={{ paddingX: 2, paddingY: 1 }}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              บันทึกจำนวนสินค้า
            </Typography>
          </Box>
        </Box>
        <Divider />
        <Grid container spacing={3} sx={{ padding: 2 }}>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>Product</p>
              <Select
                native
                value={stock.ProductID + ""}
                onChange={handleChange}
                inputProps={{
                  name: "ProductID",
                }}
              >
                <option aria-label="None" value="">
                  Please select a product.
                </option>
                {products.map((item: ProductsInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Product_name}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={12}>
            <FormControl fullWidth variant="outlined">
            <p>Quantity</p>
              <TextField
                id="Quantity"
                variant="outlined"
                type="number"
                size="medium"
                placeholder="Please enter a number."
                value={stock.Quanitiy || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>Shelf</p>
              <Select
                native
                value={stock.ShelfID + ""}
                onChange={handleChange}
                //disabled
                inputProps={{
                  name: "ShelfID",
                }}
              >
                <option aria-label="None" value="">
                  Please select shelf for product.
                </option>
                <option value={shelfs?.ID} key={shelfs?.ID}>
                  {shelfs?.Shelf_name}
                </option>
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>Lot</p>
              <Select
                native
                value={stock.ShelfID + ""}
                onChange={handleChange}
                //disabled
                inputProps={{
                  name: "LotID",
                }}
              >
                <option aria-label="None" value="">
                  Please select lot for product.
                </option>
                <option value={lots?.ID} key={lots?.ID}>
                  {lots?.Lot_number}
                </option>
              </Select>
            </FormControl>
          </Grid>
         
          <Grid item xs={12}>
            <Button
              component={RouterLink}
              to="/stocks"
              variant="contained"
              color="inherit"
            >
              กลับ
            </Button>
            <Button
              style={{ float: "right" }}
              onClick={submit}
              variant="contained"
              color="primary"
            >
              บันทึก
            </Button>
          </Grid>
        </Grid>
      </Paper>
    </Container>
  );
}

export default StockCreate;
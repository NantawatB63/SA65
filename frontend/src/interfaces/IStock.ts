import { LotsInterface } from "./ILot";
import { ShelfsInterface } from "./IShelf";
import { ProductsInterface } from "./IProduct";
import { EmployeesInterface } from "./IEmployee";

export interface StocksInterface {
    ID?: number;
    Quanitiy?: number;
    ProductID?: number;
    Product_name?: ProductsInterface;
    ShelfID?: number;
    Shelf_name?: ShelfsInterface;
    LotID?: number;
    Lot_number?: LotsInterface;
    FirstName?: EmployeesInterface;
    EmployeeID?: number;
  }
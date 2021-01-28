using System;
using System.Collections.Generic;
using System.Linq;
using System.Web;
using System.Web.UI;
using System.Web.UI.WebControls;

namespace multiArrays
{
    public partial class _default : System.Web.UI.Page
    {
        double[,] priceGrid;

        protected void Page_Load(object sender, EventArgs e)
        {
            // double[,] priceGrid = new double[3, 3];

            priceGrid = new double[3, 3];
            
            // 0 - Chicago
            // 1 - New York
            // 2 - London
           
            priceGrid[0, 1] = 350.0;
            priceGrid[0, 2] = 750.0;
            priceGrid[1, 0] = 400.0;
            priceGrid[1, 2] = 700d;
            priceGrid[2, 0] = 800d;
            priceGrid[2, 1] = 805d;

          
         
          
        }

        protected void okButton_Click(object sender, EventArgs e)
        {
            int fromCity;
            if (fromChicagoRadioButton.Checked) fromCity = 0;
            else if (fromNewYorkRadioButton.Checked) fromCity = 1;
            else fromCity = 2; 
            
            int toCity;
            if (toChicagoRadioButton.Checked) toCity = 0;
            else if (toNewYorkRadioButton.Checked) toCity = 1;
            else toCity = 2;
           
            if (fromCity == toCity) {
                resultLabel.Text = "";
                return;
            }
            else resultLabel.Text = priceGrid[fromCity, toCity].ToString();
            



            



            
        }
    }
}
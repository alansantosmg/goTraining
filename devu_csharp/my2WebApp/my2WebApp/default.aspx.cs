using System;
using System.Collections.Generic;
using System.Linq;
using System.Web;
using System.Web.UI;
using System.Web.UI.WebControls;

namespace my2WebApp
{
    public partial class _default : System.Web.UI.Page
    {
        protected void Page_Load(object sender, EventArgs e)
        {

        }

        protected void resultButton_Click(object sender, EventArgs e)
        {
            string myAge = myAgeTextBox.Text;
            string myMoney = myMoneyTextBox.Text;

            resultLabel.Text = "At " + myAge + " years old, I would expect you to have more than " + myMoney + " in the pocket.";
          
        }
    }
}
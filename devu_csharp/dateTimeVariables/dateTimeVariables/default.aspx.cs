using System;
using System.Collections.Generic;
using System.Linq;
using System.Web;
using System.Web.UI;
using System.Web.UI.WebControls;

namespace dateTimeVariables
{
    public partial class _default : System.Web.UI.Page
    {
        protected void Page_Load(object sender, EventArgs e)
        {

        }

        protected void Button1_Click(object sender, EventArgs e)
        {

            DateTime myValue = new DateTime(1973, 06, 12);
            // resultLabel.Text = myValue.ToString(); 
            resultLabel.Text = myValue.ToShortDateString();



        }
    }
}
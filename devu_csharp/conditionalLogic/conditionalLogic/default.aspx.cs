using System;
using System.Collections.Generic;
using System.Linq;
using System.Web;
using System.Web.UI;
using System.Web.UI.WebControls;

namespace conditionalLogic
{
    public partial class _default : System.Web.UI.Page
    {
        protected void Page_Load(object sender, EventArgs e)
        {

        }

        protected void Button1_Click(object sender, EventArgs e)
        {

            if (pizzaRadioButton.Checked)
            {

                resultLabel.Text = "you must be from Chicago.";

            }

            else if (saladRadioButton.Checked)
            {
                resultLabel.Text = "you must be healthy.";

            }

            else if (peanutRadioButton.Checked)
            {
                resultLabel.Text = "you must be cool.";

            }
            else
            {
                resultLabel.Text = "Please, select one option.";
            }
        }
    }
}
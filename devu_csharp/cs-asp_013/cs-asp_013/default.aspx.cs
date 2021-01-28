using System;
using System.Collections.Generic;
using System.Linq;
using System.Web;
using System.Web.UI;
using System.Web.UI.WebControls;

namespace cs_asp_013
{
    public partial class _default : System.Web.UI.Page
    {
        protected void Page_Load(object sender, EventArgs e)
        {
            comparisionTypeLabel.Text = "equal";
        }

        protected void okButton_Click(object sender, EventArgs e)
        {




            /*
            resultLabel.Text = (firstTextBox.Text == secondTextBox.Text)
                ? "yes"
                : "no";
                */

            if (myCheckBox.Checked
                && firstTextBox.Text == "Alan"
                && secondTextBox.Text == "Santos")
            {
                resultLabel.Text = "Perfect Criterea";
            }
            else
            {
                resultLabel.Text = "Try again";
            }



        }
    }
}
using System;
using System.Collections.Generic;
using System.Linq;
using System.Web;
using System.Web.UI;
using System.Web.UI.WebControls;

namespace myFirstWebApp
{
    public partial class Default : System.Web.UI.Page
    {
        protected void Page_Load(object sender, EventArgs e)
        {

        }

        protected void okButton_Click(object sender, EventArgs e)
        {

            string firstName = firstnameTextBox.Text;
            string lastName = lastNameTextBox.Text;
            resultLabel.Text = firstName + " " + lastName;

        }
    }
}
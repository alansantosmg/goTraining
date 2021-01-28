using System;
using System.Collections.Generic;
using System.Linq;
using System.Web;
using System.Web.UI;
using System.Web.UI.WebControls;

namespace formatingStrings
{
    public partial class _default : System.Web.UI.Page
    {
        protected void Page_Load(object sender, EventArgs e)
        {

        }

        protected void okButton_Click(object sender, EventArgs e)
        {
            //string result = "";
            long issn = long.Parse(ssnTextBox.Text);
            long phoneNumber = long.Parse(phoneNumberTextBox.Text);
            double salary = double.Parse(salaryTextBox.Text);
       


            string result = String.Format("Thank you for your Bussiness, {0}. <br/>" +
                "Your social Security Number is: {1: 000-000-000-00}.<br/>" +
                "Your Phone number is {2: (00) 00000-0000}. <br/>" +
                "Your Loan Date is: {3:dddd - d/M/yyyy}.<br/>" +
                " Salary: {4:C}",
                nameTextBox.Text, issn, phoneNumber,loanCalendar.SelectedDate,salary);



            resultLabel.Text = result; 

        }
    }
}
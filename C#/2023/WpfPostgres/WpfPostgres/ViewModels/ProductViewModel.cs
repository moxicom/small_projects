using System;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.ComponentModel;
using System.Linq;
using System.Runtime.CompilerServices;
using System.Text;
using System.Threading.Tasks;
using WpfPostgres.Data;
using WpfPostgres.Models;

namespace WpfPostgres.ViewModels
{
    public class ProductViewModel : INotifyPropertyChanged
    {
        private ObservableCollection<User> _users;

        public ObservableCollection<User> Users
        {
            get { return _users; }
            set
            {
                _users = value;
                OnPropertyChanged(nameof(Users));
            }
        }

        public ProductViewModel()
        {
            LoadData();
        }

        private void LoadData()
        {
            using (var dbContext = new ApplicationContext())
            {
                Users = new ObservableCollection<User>(dbContext.Users.ToList());
            }
        }

        public event PropertyChangedEventHandler? PropertyChanged;
        public void OnPropertyChanged([CallerMemberName] string prop = "")
        {
            if (PropertyChanged != null)
                PropertyChanged(this, new PropertyChangedEventArgs(prop));
        }
    }
}


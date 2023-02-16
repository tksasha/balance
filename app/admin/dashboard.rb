# frozen_string_literal: true

ActiveAdmin.register_page 'Dashboard' do
  menu priority: 1, label: proc { I18n.t('active_admin.dashboard') }

  content title: proc { I18n.t('active_admin.dashboard') } do
    h2 'Залишки'
    columns do
      Cash
        .for_dashboard
        .map do |currency, collection|
          column do
            render 'admin/dashboard/cashes/currency', currency:, collection:
          end
        end
    end

    h2 'Витрати по категоріям'
    columns do
      Item
        .for_dashboard
        .map do |currency, collection|
          column do
            h3 currency.upcase

            render 'admin/dashboard/items/currency', currency:, collection:
          end
        end
    end
  end
end

# frozen_string_literal: true

ActiveAdmin.register Cash do
  include HasCurrencyScopes

  menu priority: 4, label: proc { I18n.t('active_admin.cashes') }

  permit_params :name, :formula, :currency, :favorite, :supercategory

  filter :name

  index row_class: proc { 'cash' } do
    column :name
    column(:sum, class: 'sum') { |cash| money(cash.sum) }
    column :favorite
    actions
  end

  form do |_|
    inputs do
      input :name
      input :formula
      input :currency
      input :favorite
      input :supercategory
    end
    actions
  end
end

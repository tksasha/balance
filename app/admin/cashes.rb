# frozen_string_literal: true

ActiveAdmin.register Cash do
  include HasCurrencyScopes

  menu priority: 4, label: proc { I18n.t('active_admin.cashes') }

  permit_params :name, :formula, :currency, :favorite

  filter :name

  form do |f|
    inputs do
      input :name
      input :formula
      input :currency
    end
  end
end

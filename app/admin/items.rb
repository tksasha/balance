# frozen_string_literal: true

ActiveAdmin.register Item do
  menu priority: 2, label: proc { I18n.t('active_admin.items') }

  actions :all, except: %i[destroy new create]

  permit_params %i[date]

  includes :category

  CURRENCIES.keys.each_with_index do |currency, idx|
    scope(currency.upcase, default: idx == 0) { |scope| scope.public_send(currency)  }
  end

  filter :category, as: :select, collection: []
end

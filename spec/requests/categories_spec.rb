# frozen_string_literal: true

RSpec.describe 'Categories', type: :request do
  it_behaves_like 'new', '/categories/new.js'

  it_behaves_like 'edit', '/categories/54/edit.js' do
    before { create :category, id: 54 }
  end

  it_behaves_like 'create', '/categories.js' do
    let :valid_params do
      {
        category: {
          name: 'Drinks',
          income: true,
          visible: true,
        }
      }
    end

    let :invalid_params do
      { category: { name: '' } }
    end

    let(:success) { -> { should render_template :create } }

    let(:failure) { -> { should render_template :new } }
  end

  it_behaves_like 'update', '/categories/54.js' do
    let :valid_params do
      { category: { name: 'Drinks' } }
    end

    let :invalid_params do
      { category: { name: '' } }
    end

    before { create :category, id: 54 }

    let(:success) { -> { should render_template :update } }

    let(:failure) { -> { should render_template :edit } }
  end

  it_behaves_like 'index', '/categories.js'

  # TODO: need spec for `GET /categories.js?widget=1`
end

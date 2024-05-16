# frozen_string_literal: true

RSpec.describe 'Admin/Categories' do
  describe 'PATH update' do
    let(:category) do
      create(
        :category,
        name: 'Category #1',
        currency: 'uah',
        supercategory: 'common',
        income: false,
        visible: false
      )
    end

    before { patch "/admin/categories/#{ category.id }", params: }

    context 'with valid params' do
      let(:params) do
        {
          category: {
            name: 'Category #2',
            currency: 'usd',
            supercategory: 'children',
            income: true,
            visible: true
          }
        }
      end

      before { category.reload }

      it { is_expected.to redirect_to "/admin/categories/#{ category.id }" }

      it { expect(category.name).to eq 'Category #2' }
      it { expect(category.currency).to eq 'usd' }
      it { expect(category.supercategory).to eq 'children' }
      it { expect(category.income).to be_truthy }
      it { expect(category.visible).to be_truthy }
    end

    context 'with invalid params' do
      let(:params) { { category: { name: '' } } }

      it { is_expected.to render_template :edit }
    end
  end

  describe 'GET index.json' do
    let(:headers) { { accept: 'application/json' } }

    before do
      CURRENCIES.keys.map do |currency|
        create_list(:category, 2, currency:)
      end

      get '/admin/categories', params:, headers:
    end

    CURRENCIES.map do |currency_name, currency_id|
      context "when currency is `#{ currency_name }`" do
        let(:params) { { q: { currency_eq: currency_id } } }

        let(:categories) { Category.where(currency: currency_name).as_json(only: %i[id name]) }

        it { is_expected.to render_template :index }

        it { expect(response.parsed_body).to match_array categories }
      end
    end
  end

  describe 'POST create' do
    before { post '/admin/categories', params: }

    context 'with valid params' do
      let(:params) do
        {
          category: {
            name: 'Category #1',
            currency: 'usd',
            supercategory: 'children',
            income: true,
            visible: true
          }
        }
      end

      let(:category) { Category.last }

      it { is_expected.to redirect_to "/admin/categories/#{ category.id }" }

      it { expect(category.name).to eq 'Category #1' }
      it { expect(category.currency).to eq 'usd' }
      it { expect(category.supercategory).to eq 'children' }
      it { expect(category.income).to be_truthy }
      it { expect(category.visible).to be_truthy }
    end

    context 'with invalid params' do
      let(:params) { { category: { name: '' } } }

      it { is_expected.to render_template :new }
    end
  end
end

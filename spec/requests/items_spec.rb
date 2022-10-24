# frozen_string_literal: true

RSpec.describe 'Items', type: :request do
  %w[uah usd eur].each do |currency|
    describe 'GET index.html' do
      before { get "/#{ currency }/items" }

      it_behaves_like 'index.html'
    end

    describe 'GET index.js' do
      before do
        create_list :item, 2

        get "/#{ currency }/uah/items", xhr: true
      end

      it_behaves_like 'index.js'
    end

    describe 'POST create.js' do
      let(:category) { create(:category) }

      before { post "/#{ currency }/items", params:, xhr: true }

      context 'with valid params' do
        let(:params) do
          {
            item: {
              date: '2019-11-13',
              formula: '3.5 + 5.3',
              category_id: category.id,
              description: 'Lorem Ipsum ...'
            }
          }
        end

        let(:item) { Item.last }

        it_behaves_like 'create.js'

        it { expect(item.currency).to eq currency }

        it { expect(item.date).to eq Date.new(2019, 11, 13) }

        it { expect(item.sum).to eq 8.8 }

        it { expect(item.category).to eq category }

        it { expect(item.description).to eq 'Lorem Ipsum ...' }
      end

      context 'with invalid params' do
        let(:params) { { item: { date: '' } } }

        it_behaves_like 'new.js'
      end
    end

    describe 'GET edit.js' do
      let(:item) { create(:item, currency:) }

      before { get "/#{ currency }/items/#{ item.id }/edit", xhr: true }

      it_behaves_like 'edit.js'
    end

    describe 'PATCH update.js' do
      let(:item) { create(:item, currency:) }

      let(:category) { create(:category) }

      before { patch "/#{ currency }/items/#{ item.id }", params:, xhr: true }

      context 'with valid params' do
        let(:params) do
          {
            item: {
              date: '2019-11-13',
              formula: '1.9 + 9.2',
              category_id: category.id,
              description: 'Lorem Ipsum ...'
            }
          }
        end

        before { item.reload }

        it_behaves_like 'update.js'

        it { expect(item.currency).to eq currency }

        it { expect(item.date).to eq Date.new(2019, 11, 13) }

        it { expect(item.sum).to eq 11.1 }

        it { expect(item.category).to eq category }

        it { expect(item.description).to eq 'Lorem Ipsum ...' }
      end

      context 'with invalid params' do
        let(:params) { { item: { date: '' } } }

        it_behaves_like 'edit.js'
      end
    end

    describe 'DELETE destroy.js' do
      let(:item) { create(:item) }

      before { delete "/#{ currency }/items/#{ item.id }", xhr: true }

      it_behaves_like 'destroy.js'

      it { expect(item.reload.deleted_at).not_to be_nil }
    end
  end
end

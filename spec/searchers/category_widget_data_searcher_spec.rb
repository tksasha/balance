# frozen_string_literal: true

RSpec.describe CategoryWidgetDataSearcher do
  subject { described_class.new(currency:) }

  let(:currency) { 'usd' }

  describe '#search' do
    before do
      create(:category, :usd, :income, :visible, id: 1, name: 'Один')
      create(:category, :usd, :income, :invisible, id: 2, name: 'Два')

      create(:category, :usd, :expense, :visible, id: 3, name: 'Три')
      create(:category, :usd, :expense, :invisible, id: 4, name: 'Чотири')

      create(:category, :uah, :income, :visible, id: 5, name: "П'ять")
      create(:category, :uah, :expense, :visible, id: 6, name: 'Шість')

      create(:category, :eur, :income, :visible, id: 7, name: 'Сім')
      create(:category, :eur, :expense, :visible, id: 8, name: 'Вісім')
    end

    context 'when currency is `usd`' do
      let(:categories) do
        [
          [
            'Видатки', [
              ['Три', 3]
            ]
          ],
          [
            'Надходження',
            [
              ['Один', 1]
            ]
          ]
        ]
      end

      I18n.with_locale(:uk) do
        its(:search) { is_expected.to eq categories }
      end
    end

    context 'when currency is `uah`' do
      let(:currency) { 'uah' }

      let(:categories) do
        [
          [
            'Видатки', [
              ['Шість', 6]
            ]
          ],
          [
            'Надходження', [
              ["П'ять", 5]
            ]
          ]
        ]
      end

      I18n.with_locale(:uk) do
        its(:search) { is_expected.to eq categories }
      end
    end

    context 'when currency is `eur`' do
      let(:currency) { 'eur' }

      let(:categories) do
        [
          [
            'Видатки', [
              ['Вісім', 8]
            ]
          ],
          [
            'Надходження', [
              ['Сім', 7]
            ]
          ]
        ]
      end

      I18n.with_locale(:uk) do
        its(:search) { is_expected.to eq categories }
      end
    end
  end

  describe '.search' do
    let(:instance) { double }

    before do
      allow(described_class).to receive(:new).and_return(instance)

      allow(instance).to receive(:search)

      described_class.search(currency:)
    end

    it { expect(described_class).to have_received(:new).with(currency:) }

    it { expect(instance).to have_received(:search) }
  end
end

# frozen_string_literal: true

RSpec.describe Frontend::Reports::Consolidations do
  describe '#call' do
    subject { described_class.new(currency:, month:) }

    before do
      category_n1 = create(:category, supercategory: :common, id: 1153, name: 'Food')
      category_n2 = create(:category, supercategory: :common, id: 1154, name: 'Drinks')
      category_n3 = create(:category, supercategory: :children, id: 1155, name: 'Pocket money')
      category_n4 = create(:category, id: 1650, name: 'Salary', income: true)

      # uah, common
      create(:item, :uah, sum: 10.01, date: '2023-01-01', category: category_n1)
      create(:item, :uah, sum: 10.02, date: '2023-01-02', category: category_n1)
      create(:item, :uah, sum: 10.03, date: '2023-01-03', category: category_n2)
      create(:item, :uah, sum: 10.04, date: '2023-01-04', category: category_n2)

      # uah, children
      create(:item, :uah, sum: 10.05, date: '2023-01-05', category: category_n3)
      create(:item, :uah, sum: 10.06, date: '2023-01-06', category: category_n3)

      # uah, common, 2023-02-28
      create(:item, :uah, sum: 10.07, date: '2023-02-28', category: category_n1)
      create(:item, :uah, sum: 10.08, date: '2023-02-28', category: category_n4)

      # usd, common
      create(:item, :usd, sum: 20.00, date: '2023-05-31', category: category_n1)

      # eur, common
      create(:item, :eur, sum: 30.00, date: '2023-06-30', category: category_n1)

      # uah, deleted
      create(:item, :uah, sum: 100.0, date: '2023-01-01', deleted_at: Time.zone.today)
      # usd, deleted
      create(:item, :usd, sum: 200.0, date: '2023-01-01', deleted_at: Time.zone.today)
      # eur, deleted
      create(:item, :eur, sum: 300.0, date: '2023-01-01', deleted_at: Time.zone.today)
    end

    context 'when UAH and 2023-01' do
      let(:currency) { :uah }
      let(:month) { '2023-01' }

      # rubocop:disable RSpec/MultipleExpectations
      it do
        results = subject.call

        expect(results[0]).to contain_exactly(1, [[1, 'Drinks', 1154, 20.07], [1, 'Food', 1153, 20.03]])

        expect(results[1]).to contain_exactly(2, [[2, 'Pocket money', 1155, 20.11]])
      end
      # rubocop:enable RSpec/MultipleExpectations
    end

    context 'when UAH and 2023-02' do
      let(:currency) { :uah }
      let(:month) { '2023-02' }

      # rubocop:disable RSpec/MultipleExpectations
      it do
        results = subject.call

        expect(results[0]).to eq([0, [[0, 'Salary', 1650, 10.08]]])

        expect(results[1]).to eq([1, [[1, 'Food', 1153, 10.07]]])
      end
      # rubocop:enable RSpec/MultipleExpectations
    end

    context 'when USD and 2023-05' do
      let(:currency) { :usd }
      let(:month) { '2023-05' }

      its(:call) { is_expected.to eq([[1, [[1, 'Food', 1153, 20.00]]]]) }
    end

    context 'when EUR and 2023-06' do
      let(:currency) { :eur }
      let(:month) { '2023-06' }

      its(:call) { is_expected.to eq([[1, [[1, 'Food', 1153, 30.00]]]]) }
    end

    context 'when GBP and 2023-02' do
      let(:currency) { :gbp }
      let(:month) { '2023-02' }

      # rubocop:disable RSpec/MultipleExpectations
      it do
        results = subject.call

        expect(results[0]).to eq([0, [[0, 'Salary', 1650, 10.08]]])

        expect(results[1]).to eq([1, [[1, 'Food', 1153, 10.07]]])
      end
      # rubocop:enable RSpec/MultipleExpectations
    end

    context 'when UAH and Month.new(2023, 2)' do
      let(:currency) { :uah }
      let(:month) { Month.new(2023, 2) }

      # rubocop:disable RSpec/MultipleExpectations
      it do
        results = subject.call

        expect(results[0]).to eq([0, [[0, 'Salary', 1650, 10.08]]])

        expect(results[1]).to eq([1, [[1, 'Food', 1153, 10.07]]])
      end
      # rubocop:enable RSpec/MultipleExpectations
    end
  end

  describe '.call' do
    let(:currency) { :usd }
    let(:month) { '2023-05' }
    let(:instance) { double }

    before do
      allow(described_class).to receive(:new).and_return(instance)

      allow(instance).to receive(:call)

      described_class.call(currency:, month:)
    end

    it { expect(described_class).to have_received(:new).with(currency:, month:) }

    it { expect(instance).to have_received(:call) }
  end
end

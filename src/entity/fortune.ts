import { Entity, PrimaryGeneratedColumn, Column } from 'typeorm';

@Entity({ name: 'fortunes' })
export class Fortune {

  @PrimaryGeneratedColumn()
  id?: number;

  @Column()
  name?: string;

  @Column('text')
  content?: string;

  @Column()
  due_date?: string;

  @Column()
  created_at?: string;

  @Column()
  updated_at?: string;
}
